# InputAudio


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `Data`                                                                     | *string*                                                                   | :heavy_check_mark:                                                         | Base64 encoded audio data.                                                 |
| `Format`                                                                   | [components.Format](../../models/components/format.md)                     | :heavy_check_mark:                                                         | The format of the encoded audio data. Currently supports "wav" and "mp3".<br/> |