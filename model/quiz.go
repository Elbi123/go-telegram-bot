package model

type Quiz struct {
	Id             int
	Question       Question
	Answers        []Answer
	CorrectAnswers map[string]bool
}

type Question struct {
	Id          int
	Question    string
	Description string
	Category    string
	Difficulty  string
	Tags        []Tag
	Answer      []Answer
}

type Tag struct {
	Id   int
	Name string
}

type Answer struct {
	Id              int
	Choice          string
	ChoiceValue     bool
	CorrectAnswer   bool
	MultipleAnswers bool
	QuestionId      int
	Explanation     string
}
