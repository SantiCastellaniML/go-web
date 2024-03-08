package handlers

import (
	"encoding/json"
	"go-post/app/internal"
	"go-post/app/platform/web"
	"net/http"
)

/*
type DefaultTask struct {
	//tasks is a map that stores the tasks
	Tasks  map[int]internal.Task
	lastID int
} */
/*
type TaskRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
} */

// esto se utiliza para que se envien los datos en formato json, se busca separar las responsabilidades.
/* type TaskJSON struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
} */

/* func NewDefaultTask(tasks map[int]internal.Task, lastID int) *DefaultTask {
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
} */

func (d *DefaultTask) Create_old() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//request
		var body TaskRequestBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			//nueva forma:
			if err := web.RequestJSON(r, &body); err != nil {
				web.ResponseJSON(w, http.StatusBadRequest, map[string]any{
					"message": "invalid request body",
				})
				return
			}

			//forma vieja:
			//w.Header().Set("Content-Type", "application/json")
			//w.WriteHeader(http.StatusBadRequest)
			//a mano: w.Write([]byte(`{"message": "Invalid request Body"}`))

			//con marshal:
			//pasamos del type a bytes:
			/*
				response := map[string]any{
					"message": "Invalid request body",
				}

				bytes, _ := json.Marshal(response) //convertimos un type a bytes en formato json.
				w.Write(bytes)

			*/

			//forma más directa con encoder, escribe sobre w el dato codificado:
			/*
				json.NewEncoder(w).Encode(map[string]any{
					"message": "Invalid request body",
				})
			*/

			return
		}

		//ya parseé mis datos de JSON a una estructura de Go.
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
		//guardamos el task:
		d.Tasks[task.ID] = task

		data := TaskJSON{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
		}

		//response
		/*
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]any{
				"message": "task created",
				"data":    data,
			})
		*/

		web.ResponseJSON(w, http.StatusCreated, map[string]any{
			"message": "task created",
			"data":    data,
		})

		//w.Write([]byte(`{"id": "Task created successfully"}`))
	}
}
