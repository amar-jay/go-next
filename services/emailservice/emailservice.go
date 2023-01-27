package emailservice


// "github.com/amar-jay/go_api_boilerplate/infra/mailgunclient"
type EmailService interface {
	Welcome(toEmail string) error
	ResetPassword(toEmail string, token string) error
}

type emailService struct {
//  ctx *gin.Context
}

// NewEmailService returns a new instance of the email service
func NewEmailService() EmailService {
  return &emailService{
    //ctx: ctx
  }
}

//Welcome email
func (s *emailService) Welcome(toEmail string ) error {
  // TODO: implement email notification service
	//return errors.New("email service not implemented")
  return nil
}


// resetPassword
func (s *emailService) ResetPassword(toEmail, token string) error {
	return nil
}
