package handlers

import (
	"encoding/json"
	"go-post/app/internal"
	"go-post/app/platform/web"
	"io"
	"net/http"
)

type DefaultTask struct {
	//tasks is a map that stores the tasks
	Tasks  map[int]internal.Task
	lastID int
}

type TaskRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// esto se utiliza para que se envien los datos en formato json, se busca separar las responsabilidades.
type TaskJSON struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func NewDefaultTask(tasks map[int]internal.Task, lastID int) *DefaultTask {
	//default values
	defaultTask := make(map[int]internal.Task)
	defaultLastID := 0

	if tasks == nil {
		tasks = defaultTask
	}

	if lastID == 0 {
		lastID = defaultLastID
	}

	return &DefaultTask{
		Tasks:  defaultTask,
		lastID: defaultLastID,
	}
}

func (d *DefaultTask) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// validate token
		token := r.Header.Get("Authorization")
		if token != "12345" {
			web.ResponseJSON(w, http.StatusUnauthorized, map[string]any{
				"message": "unauthorized",
			})
			return
		}

		//request
		// - read into bytes:
		bytes, err := io.ReadAll(r.Body) //to prevent reading the body twice, we read it into bytes.
		if err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}
		//parse json to map (Dynamic)
		//a json is a map[string]any in Go.
		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		//validate required keys (if the client passed the keys):
		if _, ok := bodyMap["title"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "title is required",
			})
		}

		if _, ok := bodyMap["description"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "description is required",
			})
		}

		if _, ok := bodyMap["done"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "done is required",
			})
		}

		//parse JSON to struct (static)
		var body TaskRequestBody
		if err := json.Unmarshal(bytes, &body); err != nil {
			//nueva forma:
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "invalid request body",
			})
			return
		}

		//ya parsee mis datos de JSON a una estructura de Go.
		//ahora tengo que validar que los datos sean correctos
		//process
		d.lastID++

		// serializamos el body a una estructura de Task
		task := internal.Task{
			ID:          d.lastID,
			Title:       body.Title,
			Description: body.Description,
			Done:        body.Done,
		}

		// - validate the task:
		if task.Title == "" || len(task.Title) > 25 {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
				"message": "title is required and must be less than 25 characters",
			})

			d.lastID--
			return
		}

		//guardamos el task:
		d.Tasks[task.ID] = task

		data := TaskJSON{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
		}

		//response
		web.ResponseJSON(w, http.StatusCreated, map[string]any{
			"message": "task created",
			"data":    data,
		})
	}
}
