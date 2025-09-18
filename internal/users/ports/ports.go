package ports

import (
	"employee/internal/users/domain"
	"employee/internal/users/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	EmployeesService service.EmployeesService
}

func NewHandler(employeesService service.EmployeesService) *Handler {
	return &Handler{
		EmployeesService: employeesService,
	}
}

func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	reqBody := r.Body
	defer reqBody.Close()

	var employee domain.Employee
	if err := json.NewDecoder(reqBody).Decode(&employee); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err := h.EmployeesService.CreateEmployee(r.Context(), domain.Employee{
		FullName: employee.FullName,
		Phone:    employee.Phone,
		City:     employee.City,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
