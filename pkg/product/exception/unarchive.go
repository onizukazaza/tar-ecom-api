package exception

type UnArchive struct {}

func (e *UnArchive) Error() string {
    return "Failed to unarchive product"
}