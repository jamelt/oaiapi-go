# CreateTranslationRequest


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                      | Type                                                                                                                                                                                                                                                                                                                                                       | Required                                                                                                                                                                                                                                                                                                                                                   | Description                                                                                                                                                                                                                                                                                                                                                | Example                                                                                                                                                                                                                                                                                                                                                    |
| ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `File`                                                                                                                                                                                                                                                                                                                                                     | [components.CreateTranslationRequestFile](../../models/components/createtranslationrequestfile.md)                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                         | The audio file object (not file name) translate, in one of these formats: flac, mp3, mp4, mpeg, mpga, m4a, ogg, wav, or webm.<br/>                                                                                                                                                                                                                         |                                                                                                                                                                                                                                                                                                                                                            |
| `Model`                                                                                                                                                                                                                                                                                                                                                    | [components.CreateTranslationRequestModel](../../models/components/createtranslationrequestmodel.md)                                                                                                                                                                                                                                                       | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                         | ID of the model to use. Only `whisper-1` (which is powered by our open source Whisper V2 model) is currently available.<br/>                                                                                                                                                                                                                               | whisper-1                                                                                                                                                                                                                                                                                                                                                  |
| `Prompt`                                                                                                                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                                                                                                  | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                         | An optional text to guide the model's style or continue a previous audio segment. The [prompt](/docs/guides/speech-to-text#prompting) should be in English.<br/>                                                                                                                                                                                           |                                                                                                                                                                                                                                                                                                                                                            |
| `ResponseFormat`                                                                                                                                                                                                                                                                                                                                           | [*components.AudioResponseFormat](../../models/components/audioresponseformat.md)                                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                         | The format of the output, in one of these options: `json`, `text`, `srt`, `verbose_json`, or `vtt`.<br/>                                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                                                                                            |
| `Temperature`                                                                                                                                                                                                                                                                                                                                              | **float64*                                                                                                                                                                                                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                         | The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use [log probability](https://en.wikipedia.org/wiki/Log_probability) to automatically increase the temperature until certain thresholds are hit.<br/> |                                                                                                                                                                                                                                                                                                                                                            |