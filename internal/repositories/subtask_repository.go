package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/go-redis/redis/v8"
	"github.com/nunthapongnp/time-doo-api/internal/models"
)

const (
	subtaskCacheDuration = 5 * time.Minute
)

type SubTaskRepository struct {
	fs         *firestore.Client
	redis      *redis.Client
	collection string
}

func NewSubTaskRepository(fs *firestore.Client, redis *redis.Client) *SubTaskRepository {
	return &SubTaskRepository{
		fs:         fs,
		redis:      redis,
		collection: "subtasks",
	}
}

// Helper function to generate cache keys
func (r *SubTaskRepository) taskCacheKey(id string) string {
	return fmt.Sprintf("task:%s", id)
}

func (r *SubTaskRepository) subtaskCacheKey(taskID, subtaskID string) string {
	return fmt.Sprintf("task:%s:subtask:%s", taskID, subtaskID)
}

// Subtask CRUD with Redis integration
func (r *SubTaskRepository) CreateSubtask(ctx context.Context, taskID string, subtask *models.Subtask) (string, error) {
	// Firestore operation
	docRef, _, err := r.fs.Collection(r.collection).Doc(taskID).
		Collection("subtasks").Add(ctx, subtask)
	if err != nil {
		return "", err
	}
	subtaskID := docRef.ID

	// Invalidate cache
	if err := r.redis.Del(ctx, r.taskCacheKey(taskID)).Err(); err != nil {
		fmt.Printf("Cache delete error: %v\n", err)
	}

	return subtaskID, nil
}

func (r *SubTaskRepository) GetSubtask(ctx context.Context, taskID, subtaskID string) (*models.Subtask, error) {
	// Try cache first
	cachedSubtask, err := r.getCachedSubtask(ctx, taskID, subtaskID)
	if err == nil {
		return cachedSubtask, nil
	}

	// Firestore operation
	docSnap, err := r.fs.Collection(r.collection).Doc(taskID).
		Collection("subtasks").Doc(subtaskID).Get(ctx)
	if err != nil {
		return nil, err
	}

	subtask := new(models.Subtask)
	if err := docSnap.DataTo(subtask); err != nil {
		return nil, err
	}
	subtask.ID = docSnap.Ref.ID
	subtask.TaskID = taskID

	// Cache the subtask
	if err := r.cacheSubtask(ctx, taskID, subtaskID, subtask); err != nil {
		fmt.Printf("Cache set error: %v\n", err)
	}

	return subtask, nil
}

func (r *SubTaskRepository) UpdateSubtask(ctx context.Context, taskID, subtaskID string, subtask *models.Subtask) error {
	// Firestore operation
	_, err := r.fs.Collection(r.collection).Doc(taskID).
		Collection("subtasks").Doc(subtaskID).Update(ctx, subtask.ToFirestoreUpdate())
	if err != nil {
		return err
	}

	// Update cache
	if err := r.cacheSubtask(ctx, taskID, subtaskID, subtask); err != nil {
		fmt.Printf("Cache update error: %v\n", err)
	}

	// Invalidate parent task cache
	if err := r.redis.Del(ctx, r.taskCacheKey(taskID)).Err(); err != nil {
		fmt.Printf("Cache invalidation error: %v\n", err)
	}

	return nil
}

func (r *SubTaskRepository) DeleteSubtask(ctx context.Context, taskID, subtaskID string) error {
	// Firestore operation
	_, err := r.fs.Collection(r.collection).Doc(taskID).
		Collection("subtasks").Doc(subtaskID).Delete(ctx)
	if err != nil {
		return err
	}

	// Invalidate caches
	if err := r.redis.Del(
		ctx,
		r.subtaskCacheKey(taskID, subtaskID),
		r.taskCacheKey(taskID),
	).Err(); err != nil {
		fmt.Printf("Cache delete error: %v\n", err)
	}

	return nil
}

func (r *SubTaskRepository) cacheSubtask(ctx context.Context, taskID, subtaskID string, subtask *models.Subtask) error {
	serialized, err := json.Marshal(subtask)
	if err != nil {
		return err
	}
	return r.redis.Set(
		ctx,
		r.subtaskCacheKey(taskID, subtaskID),
		serialized,
		subtaskCacheDuration,
	).Err()
}

func (r *SubTaskRepository) getCachedSubtask(ctx context.Context, taskID, subtaskID string) (*models.Subtask, error) {
	val, err := r.redis.Get(ctx, r.subtaskCacheKey(taskID, subtaskID)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("not found in cache")
	}
	if err != nil {
		return nil, err
	}

	subtask := new(models.Subtask)
	if err := json.Unmarshal([]byte(val), subtask); err != nil {
		return nil, err
	}
	return subtask, nil
}
