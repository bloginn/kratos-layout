package biz

import "github.com/google/wire"

// ProviderBizSet is biz providers.
var ProviderBizSet = wire.NewSet(NewOpenUseCase)
