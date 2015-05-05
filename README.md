## UUID4

Random UUIDs done simply.

    import "github.com/jamesandariese/uuid4"

    u := uuid4.NewUUID()
    u.HexString()
    u.DashString()
    u.Bytes()

or the most common use:

    uuid4.DashString()
