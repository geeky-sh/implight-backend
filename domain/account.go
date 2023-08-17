package domain

type GoogleCallback struct {
	Credential string `json:"credential"`
	GCSRFToken string `json:"g_csrf_token"`
}
