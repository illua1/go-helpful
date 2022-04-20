package helpful_map

type Map[K comparable, T any][]MapKey[K, T]

type MapKey[K comparable, T any]struct{
  key K
  value T
}

func NewMap[K comparable, T any]()Map[K, T]{
  return Map[K, T]{}
}

func(map_ *Map[K, T])Find(key K)(ret T, ok bool){
  for i := range *map_ {
    if (*map_)[i].key == key {
      ret = (*map_)[i].value
      ok = true
      return
    }
  }
  ok = false
  return
}

func(map_ *Map[K, T])Set(key K, value T){
  for i := range *map_ {
    if (*map_)[i].key == key {
      (*map_)[i].value = value
      return
    }
  }
  (*map_) = append((*map_), MapKey[K, T]{key, value})
  return
}