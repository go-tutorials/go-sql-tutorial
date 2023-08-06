package filter

type UserFilter struct {
	Id        string `yaml:"id" mapstructure:"id" json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"id" validate:"required,max=40" operator:"="`
	Username  string `yaml:"username" mapstructure:"username" json:"username" gorm:"column:username" bson:"username" dynamodbav:"username" firestore:"username" validate:"required,username,max=100"`
	Email     string `yaml:"email" mapstructure:"email" json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" validate:"email,max=100"`
	Phone     string `yaml:"phone" mapstructure:"phone" json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" validate:"required,phone,max=18" operator:"like"`
	PageIndex int64  `yaml:"page_index" mapstructure:"page_index" json:"pageIndex,omitempty" gorm:"column:page_index" bson:"pageIndex,omitempty" dynamodbav:"pageIndex,omitempty" firestore:"pageIndex,omitempty"`
	PageSize  int64  `yaml:"page_size" mapstructure:"page_size" json:"pageSize,omitempty" gorm:"column:page_size" bson:"pageSize,omitempty" dynamodbav:"pageSize,omitempty" firestore:"pageSize,omitempty"`
}
