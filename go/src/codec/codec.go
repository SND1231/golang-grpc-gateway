package codes

type Code string

const (
        OK Code = "OK"

        InvalidSyntax  Code = "invalid_syntax"
        BadParams      Code = "bad_params"
        EmptyBody      Code = "empty_body"
        InvalidRequest Code = "invalid_request"

        Unauthorized Code = "unauthorized"
        NotFound Code = "not_found"
        Database    Code = "database_error"
        Forbidden Code = "forbidden"

        Unknown Code = "unknown"
)
