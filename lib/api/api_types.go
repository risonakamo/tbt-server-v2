// types that support api

package timeblock_api

// API reques to change a block's title
type TitleChangeReq struct {
    BlockId string
    NewTitle string
}