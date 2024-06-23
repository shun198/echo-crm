package serializers

type (
	UpdateUserPassword struct {
		CurrentPassword string `json:"currentPassword" validate:"required,min=1"`
		NewPassword     string `json:"newPassword" validate:"required,min=1,max=64"`
		ConfirmPassword string `json:"confirmPassword" validate:"required,min=1,max=64"`
	}

	LoginCredentials struct {
		EmployeeNumber string `json:"employee_number" validate:"required"`
		Password       string `json:"password" validate:"required"`
	}

	UserInfo struct {
		Name           string `json:"name"`
		EmployeeNumber string `json:"employeeNumber"`
		Role           string `json:"role"`
		Message        string `json:"msg,omitempty"`
	}

	SendResetPasswordEmail struct {
		EmployeeNumber string `json:"employeeNumber" validate:"required"`
	}

	ConfirmPassword struct {
		Token           string `json:"token"`
		NewPassword     string `json:"newPassword" validate:"required,max=64"`
		ConfirmPassword string `json:"confirmPassword" validate:"required,max=64"`
	}

	ChangeUserDetails struct {
		Name  string `json:"name" validate:"required"`
		Role  string `json:"role" validate:"required"`
		Email string `json:"email" validate:"required"`
	}

	ErrorMessage struct {
		Name  string `json:"name"`
		Error string `json:"error"`
	}

	CheckResponse struct {
		Check bool `json:"check"`
	}

	SignUp struct {
		EmployeeNumber string `json:"employee_number" validate:"required"`
		Name           string `json:"name" validate:"required"`
		Role           string `json:"role" validate:"required"`
		Email          string `json:"email" validate:"required"`
	}

	CheckToken struct {
		Token string `json:"token"`
	}
)
