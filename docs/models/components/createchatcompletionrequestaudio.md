# CreateChatCompletionRequestAudio

Parameters for audio output. Required when audio output is requested with
`modalities: ["audio"]`. [Learn more](/docs/guides/audio).



## Fields

| Field                                                                                                                                 | Type                                                                                                                                  | Required                                                                                                                              | Description                                                                                                                           |
| ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| `Voice`                                                                                                                               | [components.CreateChatCompletionRequestVoice](../../models/components/createchatcompletionrequestvoice.md)                            | :heavy_check_mark:                                                                                                                    | The voice the model uses to respond. Supported voices are `alloy`,<br/>`ash`, `ballad`, `coral`, `echo`, `sage`, `shimmer`, and `verse`.<br/> |
| `Format`                                                                                                                              | [components.CreateChatCompletionRequestFormat](../../models/components/createchatcompletionrequestformat.md)                          | :heavy_check_mark:                                                                                                                    | Specifies the output audio format. Must be one of `wav`, `mp3`, `flac`,<br/>`opus`, or `pcm16`.<br/>                                  |