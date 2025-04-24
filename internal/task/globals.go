package task

var (
    DefaultIDGenerator IDGenerator = &TimeBasedIDGenerator{}
    DefaultStorage Storage = nil
)