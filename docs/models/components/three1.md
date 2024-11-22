# Three1

An object describing an image to classify.


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `Type`                                                                 | [components.ThreeType](../../models/components/threetype.md)           | :heavy_check_mark:                                                     | Always `image_url`.                                                    |
| `ImageURL`                                                             | [components.ThreeImageURL](../../models/components/threeimageurl.md)   | :heavy_check_mark:                                                     | Contains either an image URL or a data URL for a base64 encoded image. |