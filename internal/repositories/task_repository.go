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
	taskCacheDuration = 5 * time.Minute
)

type TaskRepository struct {
	fs         *firestore.Client
	redis      *redis.Client
	collection string
}

func NewTaskRepository(fs *firestore.Client, redis *redis.Client) *TaskRepository {
	return &TaskRepository{
		fs:         fs,
		redis:      redis,
		collection: "tasks",
	}
}

func (r *TaskRepository) taskCacheKey(id string) string {
	return fmt.Sprintf("task:%s", id)
}

func (r *TaskRepository) CreateTask(ctx context.Context, task *models.Task) (string, error) {
	docRef, _, err := r.fs.Collection(r.collection).Add(ctx, task)
	if err != nil {
		return "", err
	}
	taskID := docRef.ID
	return taskID, nil
}

func (r *TaskRepository) GetTask(ctx context.Context, id string) (*models.Task, error) {
	// Try cache first
	cachedTask, err := r.getCachedTask(ctx, id)
	if err == nil {
		return cachedTask, nil
	}

	// Fallback to Firestore
	docSnap, err := r.fs.Collection(r.collection).Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	task := new(models.Task)
	if err := docSnap.DataTo(task); err != nil {
		return nil, err
	}
	task.ID = docSnap.Ref.ID

	// Get subtasks
	subtasks, err := r.fs.Collection(r.collection).Doc(id).
		Collection("subtasks").Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	for _, doc := range subtasks {
		var subtask models.Subtask
		doc.DataTo(&subtask)
		subtask.ID = doc.Ref.ID
		subtask.TaskID = id
		task.Subtasks = append(task.Subtasks, subtask)
	}

	// Cache the result
	if err := r.cacheTask(ctx, id, task); err != nil {
		fmt.Printf("Cache set error: %v\n", err)
	}

	return task, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, id string, task *models.Task) error {
	// Firestore operation
	_, err := r.fs.Collection(r.collection).Doc(id).Update(ctx, task.ToFirestoreUpdate())
	if err != nil {
		return err
	}

	// Invalidate cache
	if err := r.redis.Del(ctx, r.taskCacheKey(id)).Err(); err != nil {
		fmt.Printf("Cache delete error: %v\n", err)
	}

	return nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id string) error {
	// Firestore operation
	_, err := r.fs.Collection(r.collection).Doc(id).Delete(ctx)
	if err != nil {
		return err
	}

	// Invalidate cache
	if err := r.redis.Del(ctx, r.taskCacheKey(id)).Err(); err != nil {
		fmt.Printf("Cache delete error: %v\n", err)
	}

	return nil
}

// Cache helper methods
func (r *TaskRepository) cacheTask(ctx context.Context, id string, task *models.Task) error {
	serialized, err := json.Marshal(task)
	if err != nil {
		return err
	}
	return r.redis.Set(ctx, r.taskCacheKey(id), serialized, taskCacheDuration).Err()
}

func (r *TaskRepository) getCachedTask(ctx context.Context, id string) (*models.Task, error) {
	val, err := r.redis.Get(ctx, r.taskCacheKey(id)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, fmt.Errorf("not found in cache")
	}
	if err != nil {
		return nil, err
	}

	task := new(models.Task)
	if err := json.Unmarshal([]byte(val), task); err != nil {
		return nil, err
	}
	return task, nil
}
