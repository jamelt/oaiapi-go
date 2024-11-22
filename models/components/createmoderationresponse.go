// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Categories - A list of the categories, and whether they are flagged or not.
type Categories struct {
	// Content that expresses, incites, or promotes hate based on race, gender, ethnicity, religion, nationality, sexual orientation, disability status, or caste. Hateful content aimed at non-protected groups (e.g., chess players) is harassment.
	Hate bool `json:"hate"`
	// Hateful content that also includes violence or serious harm towards the targeted group based on race, gender, ethnicity, religion, nationality, sexual orientation, disability status, or caste.
	HateThreatening bool `json:"hate/threatening"`
	// Content that expresses, incites, or promotes harassing language towards any target.
	Harassment bool `json:"harassment"`
	// Harassment content that also includes violence or serious harm towards any target.
	HarassmentThreatening bool `json:"harassment/threatening"`
	// Content that includes instructions or advice that facilitate the planning or execution of wrongdoing, or that gives advice or instruction on how to commit illicit acts. For example, "how to shoplift" would fit this category.
	Illicit bool `json:"illicit"`
	// Content that includes instructions or advice that facilitate the planning or execution of wrongdoing that also includes violence, or that gives advice or instruction on the procurement of any weapon.
	IllicitViolent bool `json:"illicit/violent"`
	// Content that promotes, encourages, or depicts acts of self-harm, such as suicide, cutting, and eating disorders.
	SelfHarm bool `json:"self-harm"`
	// Content where the speaker expresses that they are engaging or intend to engage in acts of self-harm, such as suicide, cutting, and eating disorders.
	SelfHarmIntent bool `json:"self-harm/intent"`
	// Content that encourages performing acts of self-harm, such as suicide, cutting, and eating disorders, or that gives instructions or advice on how to commit such acts.
	SelfHarmInstructions bool `json:"self-harm/instructions"`
	// Content meant to arouse sexual excitement, such as the description of sexual activity, or that promotes sexual services (excluding sex education and wellness).
	Sexual bool `json:"sexual"`
	// Sexual content that includes an individual who is under 18 years old.
	SexualMinors bool `json:"sexual/minors"`
	// Content that depicts death, violence, or physical injury.
	Violence bool `json:"violence"`
	// Content that depicts death, violence, or physical injury in graphic detail.
	ViolenceGraphic bool `json:"violence/graphic"`
}

func (o *Categories) GetHate() bool {
	if o == nil {
		return false
	}
	return o.Hate
}

func (o *Categories) GetHateThreatening() bool {
	if o == nil {
		return false
	}
	return o.HateThreatening
}

func (o *Categories) GetHarassment() bool {
	if o == nil {
		return false
	}
	return o.Harassment
}

func (o *Categories) GetHarassmentThreatening() bool {
	if o == nil {
		return false
	}
	return o.HarassmentThreatening
}

func (o *Categories) GetIllicit() bool {
	if o == nil {
		return false
	}
	return o.Illicit
}

func (o *Categories) GetIllicitViolent() bool {
	if o == nil {
		return false
	}
	return o.IllicitViolent
}

func (o *Categories) GetSelfHarm() bool {
	if o == nil {
		return false
	}
	return o.SelfHarm
}

func (o *Categories) GetSelfHarmIntent() bool {
	if o == nil {
		return false
	}
	return o.SelfHarmIntent
}

func (o *Categories) GetSelfHarmInstructions() bool {
	if o == nil {
		return false
	}
	return o.SelfHarmInstructions
}

func (o *Categories) GetSexual() bool {
	if o == nil {
		return false
	}
	return o.Sexual
}

func (o *Categories) GetSexualMinors() bool {
	if o == nil {
		return false
	}
	return o.SexualMinors
}

func (o *Categories) GetViolence() bool {
	if o == nil {
		return false
	}
	return o.Violence
}

func (o *Categories) GetViolenceGraphic() bool {
	if o == nil {
		return false
	}
	return o.ViolenceGraphic
}

// CategoryScores - A list of the categories along with their scores as predicted by model.
type CategoryScores struct {
	// The score for the category 'hate'.
	Hate float64 `json:"hate"`
	// The score for the category 'hate/threatening'.
	HateThreatening float64 `json:"hate/threatening"`
	// The score for the category 'harassment'.
	Harassment float64 `json:"harassment"`
	// The score for the category 'harassment/threatening'.
	HarassmentThreatening float64 `json:"harassment/threatening"`
	// The score for the category 'illicit'.
	Illicit float64 `json:"illicit"`
	// The score for the category 'illicit/violent'.
	IllicitViolent float64 `json:"illicit/violent"`
	// The score for the category 'self-harm'.
	SelfHarm float64 `json:"self-harm"`
	// The score for the category 'self-harm/intent'.
	SelfHarmIntent float64 `json:"self-harm/intent"`
	// The score for the category 'self-harm/instructions'.
	SelfHarmInstructions float64 `json:"self-harm/instructions"`
	// The score for the category 'sexual'.
	Sexual float64 `json:"sexual"`
	// The score for the category 'sexual/minors'.
	SexualMinors float64 `json:"sexual/minors"`
	// The score for the category 'violence'.
	Violence float64 `json:"violence"`
	// The score for the category 'violence/graphic'.
	ViolenceGraphic float64 `json:"violence/graphic"`
}

func (o *CategoryScores) GetHate() float64 {
	if o == nil {
		return 0.0
	}
	return o.Hate
}

func (o *CategoryScores) GetHateThreatening() float64 {
	if o == nil {
		return 0.0
	}
	return o.HateThreatening
}

func (o *CategoryScores) GetHarassment() float64 {
	if o == nil {
		return 0.0
	}
	return o.Harassment
}

func (o *CategoryScores) GetHarassmentThreatening() float64 {
	if o == nil {
		return 0.0
	}
	return o.HarassmentThreatening
}

func (o *CategoryScores) GetIllicit() float64 {
	if o == nil {
		return 0.0
	}
	return o.Illicit
}

func (o *CategoryScores) GetIllicitViolent() float64 {
	if o == nil {
		return 0.0
	}
	return o.IllicitViolent
}

func (o *CategoryScores) GetSelfHarm() float64 {
	if o == nil {
		return 0.0
	}
	return o.SelfHarm
}

func (o *CategoryScores) GetSelfHarmIntent() float64 {
	if o == nil {
		return 0.0
	}
	return o.SelfHarmIntent
}

func (o *CategoryScores) GetSelfHarmInstructions() float64 {
	if o == nil {
		return 0.0
	}
	return o.SelfHarmInstructions
}

func (o *CategoryScores) GetSexual() float64 {
	if o == nil {
		return 0.0
	}
	return o.Sexual
}

func (o *CategoryScores) GetSexualMinors() float64 {
	if o == nil {
		return 0.0
	}
	return o.SexualMinors
}

func (o *CategoryScores) GetViolence() float64 {
	if o == nil {
		return 0.0
	}
	return o.Violence
}

func (o *CategoryScores) GetViolenceGraphic() float64 {
	if o == nil {
		return 0.0
	}
	return o.ViolenceGraphic
}

type Hate string

const (
	HateText Hate = "text"
)

func (e Hate) ToPointer() *Hate {
	return &e
}
func (e *Hate) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = Hate(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Hate: %v", v)
	}
}

type HateThreatening string

const (
	HateThreateningText HateThreatening = "text"
)

func (e HateThreatening) ToPointer() *HateThreatening {
	return &e
}
func (e *HateThreatening) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = HateThreatening(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HateThreatening: %v", v)
	}
}

type Harassment string

const (
	HarassmentText Harassment = "text"
)

func (e Harassment) ToPointer() *Harassment {
	return &e
}
func (e *Harassment) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = Harassment(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Harassment: %v", v)
	}
}

type HarassmentThreatening string

const (
	HarassmentThreateningText HarassmentThreatening = "text"
)

func (e HarassmentThreatening) ToPointer() *HarassmentThreatening {
	return &e
}
func (e *HarassmentThreatening) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = HarassmentThreatening(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HarassmentThreatening: %v", v)
	}
}

type Illicit string

const (
	IllicitText Illicit = "text"
)

func (e Illicit) ToPointer() *Illicit {
	return &e
}
func (e *Illicit) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = Illicit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Illicit: %v", v)
	}
}

type IllicitViolent string

const (
	IllicitViolentText IllicitViolent = "text"
)

func (e IllicitViolent) ToPointer() *IllicitViolent {
	return &e
}
func (e *IllicitViolent) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = IllicitViolent(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IllicitViolent: %v", v)
	}
}

type SelfHarm string

const (
	SelfHarmText  SelfHarm = "text"
	SelfHarmImage SelfHarm = "image"
)

func (e SelfHarm) ToPointer() *SelfHarm {
	return &e
}
func (e *SelfHarm) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "image":
		*e = SelfHarm(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SelfHarm: %v", v)
	}
}

type SelfHarmIntent string

const (
	SelfHarmIntentText  SelfHarmIntent = "text"
	SelfHarmIntentImage SelfHarmIntent = "image"
)

func (e SelfHarmIntent) ToPointer() *SelfHarmIntent {
	return &e
}
func (e *SelfHarmIntent) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "image":
		*e = SelfHarmIntent(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SelfHarmIntent: %v", v)
	}
}

type SelfHarmInstructions string

const (
	SelfHarmInstructionsText  SelfHarmInstructions = "text"
	SelfHarmInstructionsImage SelfHarmInstructions = "image"
)

func (e SelfHarmInstructions) ToPointer() *SelfHarmInstructions {
	return &e
}
func (e *SelfHarmInstructions) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "image":
		*e = SelfHarmInstructions(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SelfHarmInstructions: %v", v)
	}
}

type Sexual string

const (
	SexualText  Sexual = "text"
	SexualImage Sexual = "image"
)

func (e Sexual) ToPointer() *Sexual {
	return &e
}
func (e *Sexual) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "image":
		*e = Sexual(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sexual: %v", v)
	}
}

type SexualMinors string

const (
	SexualMinorsText SexualMinors = "text"
)

func (e SexualMinors) ToPointer() *SexualMinors {
	return &e
}
func (e *SexualMinors) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		*e = SexualMinors(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SexualMinors: %v", v)
	}
}

type Violence string

const (
	ViolenceText  Violence = "text"
	ViolenceImage Violence = "image"
)

func (e Violence) ToPointer() *Violence {
	return &e
}
func (e *Violence) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "image":
		*e = Violence(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Violence: %v", v)
	}
}

type ViolenceGraphic string

const (
	ViolenceGraphicText  ViolenceGraphic = "text"
	ViolenceGraphicImage ViolenceGraphic = "image"
)

func (e ViolenceGraphic) ToPointer() *ViolenceGraphic {
	return &e
}
func (e *ViolenceGraphic) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text":
		fallthrough
	case "image":
		*e = ViolenceGraphic(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ViolenceGraphic: %v", v)
	}
}

// CategoryAppliedInputTypes - A list of the categories along with the input type(s) that the score applies to.
type CategoryAppliedInputTypes struct {
	// The applied input type(s) for the category 'hate'.
	Hate []Hate `json:"hate"`
	// The applied input type(s) for the category 'hate/threatening'.
	HateThreatening []HateThreatening `json:"hate/threatening"`
	// The applied input type(s) for the category 'harassment'.
	Harassment []Harassment `json:"harassment"`
	// The applied input type(s) for the category 'harassment/threatening'.
	HarassmentThreatening []HarassmentThreatening `json:"harassment/threatening"`
	// The applied input type(s) for the category 'illicit'.
	Illicit []Illicit `json:"illicit"`
	// The applied input type(s) for the category 'illicit/violent'.
	IllicitViolent []IllicitViolent `json:"illicit/violent"`
	// The applied input type(s) for the category 'self-harm'.
	SelfHarm []SelfHarm `json:"self-harm"`
	// The applied input type(s) for the category 'self-harm/intent'.
	SelfHarmIntent []SelfHarmIntent `json:"self-harm/intent"`
	// The applied input type(s) for the category 'self-harm/instructions'.
	SelfHarmInstructions []SelfHarmInstructions `json:"self-harm/instructions"`
	// The applied input type(s) for the category 'sexual'.
	Sexual []Sexual `json:"sexual"`
	// The applied input type(s) for the category 'sexual/minors'.
	SexualMinors []SexualMinors `json:"sexual/minors"`
	// The applied input type(s) for the category 'violence'.
	Violence []Violence `json:"violence"`
	// The applied input type(s) for the category 'violence/graphic'.
	ViolenceGraphic []ViolenceGraphic `json:"violence/graphic"`
}

func (o *CategoryAppliedInputTypes) GetHate() []Hate {
	if o == nil {
		return []Hate{}
	}
	return o.Hate
}

func (o *CategoryAppliedInputTypes) GetHateThreatening() []HateThreatening {
	if o == nil {
		return []HateThreatening{}
	}
	return o.HateThreatening
}

func (o *CategoryAppliedInputTypes) GetHarassment() []Harassment {
	if o == nil {
		return []Harassment{}
	}
	return o.Harassment
}

func (o *CategoryAppliedInputTypes) GetHarassmentThreatening() []HarassmentThreatening {
	if o == nil {
		return []HarassmentThreatening{}
	}
	return o.HarassmentThreatening
}

func (o *CategoryAppliedInputTypes) GetIllicit() []Illicit {
	if o == nil {
		return []Illicit{}
	}
	return o.Illicit
}

func (o *CategoryAppliedInputTypes) GetIllicitViolent() []IllicitViolent {
	if o == nil {
		return []IllicitViolent{}
	}
	return o.IllicitViolent
}

func (o *CategoryAppliedInputTypes) GetSelfHarm() []SelfHarm {
	if o == nil {
		return []SelfHarm{}
	}
	return o.SelfHarm
}

func (o *CategoryAppliedInputTypes) GetSelfHarmIntent() []SelfHarmIntent {
	if o == nil {
		return []SelfHarmIntent{}
	}
	return o.SelfHarmIntent
}

func (o *CategoryAppliedInputTypes) GetSelfHarmInstructions() []SelfHarmInstructions {
	if o == nil {
		return []SelfHarmInstructions{}
	}
	return o.SelfHarmInstructions
}

func (o *CategoryAppliedInputTypes) GetSexual() []Sexual {
	if o == nil {
		return []Sexual{}
	}
	return o.Sexual
}

func (o *CategoryAppliedInputTypes) GetSexualMinors() []SexualMinors {
	if o == nil {
		return []SexualMinors{}
	}
	return o.SexualMinors
}

func (o *CategoryAppliedInputTypes) GetViolence() []Violence {
	if o == nil {
		return []Violence{}
	}
	return o.Violence
}

func (o *CategoryAppliedInputTypes) GetViolenceGraphic() []ViolenceGraphic {
	if o == nil {
		return []ViolenceGraphic{}
	}
	return o.ViolenceGraphic
}

type Results struct {
	// Whether any of the below categories are flagged.
	Flagged bool `json:"flagged"`
	// A list of the categories, and whether they are flagged or not.
	Categories Categories `json:"categories"`
	// A list of the categories along with their scores as predicted by model.
	CategoryScores CategoryScores `json:"category_scores"`
	// A list of the categories along with the input type(s) that the score applies to.
	CategoryAppliedInputTypes CategoryAppliedInputTypes `json:"category_applied_input_types"`
}

func (o *Results) GetFlagged() bool {
	if o == nil {
		return false
	}
	return o.Flagged
}

func (o *Results) GetCategories() Categories {
	if o == nil {
		return Categories{}
	}
	return o.Categories
}

func (o *Results) GetCategoryScores() CategoryScores {
	if o == nil {
		return CategoryScores{}
	}
	return o.CategoryScores
}

func (o *Results) GetCategoryAppliedInputTypes() CategoryAppliedInputTypes {
	if o == nil {
		return CategoryAppliedInputTypes{}
	}
	return o.CategoryAppliedInputTypes
}

// CreateModerationResponse - Represents if a given text input is potentially harmful.
type CreateModerationResponse struct {
	// The unique identifier for the moderation request.
	ID string `json:"id"`
	// The model used to generate the moderation results.
	Model string `json:"model"`
	// A list of moderation objects.
	Results []Results `json:"results"`
}

func (o *CreateModerationResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateModerationResponse) GetModel() string {
	if o == nil {
		return ""
	}
	return o.Model
}

func (o *CreateModerationResponse) GetResults() []Results {
	if o == nil {
		return []Results{}
	}
	return o.Results
}