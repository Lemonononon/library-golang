package define

const LibraryResponse = "T_CYCLE_RESPONSE"
const LibraryUserID = "USER_ID"
const LibraryUserRole = "USER_ROLE"
const LibraryToken = "LIBRARY_AUTH_BEARER"

type CardIdentity int

const (
	CardTeacher int = iota
	CardStudent
)
