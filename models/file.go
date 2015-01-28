package models

type VideoFile interface {
  IsValidMKV() bool
  IsMKV() bool
  HasH264Stream() bool
  HasAC3Stream() bool
}

type File struct {
  Streams     []Stream    `json:"streams"`
  Format      *Format     `json:"format"`
}

func (f *File) IsValidMKV() bool {
  return f.IsMKV() && f.HasH264Stream() && f.HasAC3Stream()
}

func (f *File) IsMKV() bool {
  if f.Format == nil {
    return false
  }
  return (f.Format.FormatName == "matroska,webm")
}

func (f *File) HasH264Stream() bool {
  for _, stream := range f.Streams {
    if stream.CodecName == "h264" {
      return true
    }
  }
  return false
}

func (f *File) HasAC3Stream() bool {
  for _, stream := range f.Streams {
    if stream.CodecName == "ac3" {
      return true
    }
  }
  return false
}