package datacache

import "goMediatools/model"

const CN = "zh-CN"
const US = "en-US"

var LocalCache = make(map[string]model.Movie)
var TMDBCache = make(map[int]model.BaseMovie)
