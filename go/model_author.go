/*
 * Books API
 *
 * This web service offers information on books
 *
 * API version: 0.1.9
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Author struct {

	AuthorId string `json:"authorId,omitempty"`

	Name string `json:"name,omitempty"`

	Nationality string `json:"nationality,omitempty"`

	Birth string `json:"birth,omitempty"`

	Genere string `json:"genere,omitempty"`

	Books *Book `json:"books,omitempty"`
}
