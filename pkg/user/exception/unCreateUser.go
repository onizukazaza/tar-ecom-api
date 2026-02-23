package exception


type UnCreateUser struct {

}

func (e *UnCreateUser) Error() string {
    return "User creation failed"
}