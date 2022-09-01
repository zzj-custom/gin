package email_event

func (e *EmailInfo) GetMailTo() []string {
	return e.MailTo
}

func (e *EmailInfo) GetSubject() string {
	return e.Subject
}

func (e *EmailInfo) GetText() string {
	return e.Text
}

func (e *EmailInfo) GetHtml() string {
	return e.Html
}

func (e *EmailInfo) GetAttachFile() []string {
	return e.AttachFile
}

func (e *EmailInfo) GetCarbonCopy() []string {
	return e.CarbonCopy
}

func (e *EmailInfo) SetMailTo(mailTo []string) {
	e.MailTo = mailTo
}

func (e *EmailInfo) SetSubject(subject string) {
	e.Subject = subject
}

func (e *EmailInfo) SetText(text string) {
	e.Text = text
}

func (e *EmailInfo) SetHtml(html string) {
	e.Html = html
}

func (e *EmailInfo) SetAttachFile(attachFile []string) {
	e.AttachFile = attachFile
}

func (e *EmailInfo) SetCarbonCopy(carbonCopy []string) {
	e.CarbonCopy = carbonCopy
}
