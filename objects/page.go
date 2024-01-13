package objects

// WikiPage struct. https://dev.vk.com/ru/reference/objects/wiki-page
type WikiPage struct {
	ID                       int    `json:"id"`       // Page ID.
	GroupID                  int    `json:"group_id"` // Community ID.
	CreatorID                int    `json:"creator_id"`
	Title                    string `json:"title"`
	CurrentUserCanEdit       int    `json:"current_user_can_edit"`        // 1, если текущий пользователь может редактировать текст вики-страницы, иначе — 0
	CurrentUserCanEditAccess int    `json:"current_user_can_edit_access"` // 1, если текущий пользователь может изменять права доступа на вики-страницу, иначе — 0.
	WhoCanEdit               int    `json:"who_can_edit"`
	WhoCanView               int    `json:"who_can_view"`
	Edited                   int    `json:"edited"`
	Created                  int    `json:"created"`
	EditorID                 int    `json:"editor_id"`
	Views                    int    `json:"views"`
	Parent                   string `json:"parent"`
	Parent2                  string `json:"parent2"`
	Source                   string `json:"source"`
	HTML                     string `json:"html"`
	ViewURL                  string `json:"view_url"`
}
