package blocks

// import (
// 	"database/sql"
// 	_ "github.com/Go-SQL-Driver/MySQL"
// 	"log"
// )

// type ToolsBlock struct {
// 	Skills []Skill
// 	Tools  []Tool
// }

// type Skill struct {
// }

// type Tool struct {
// }

// func (t ToolsBlock) RetrieveById(db *sql.DB, id string) ToolsBlock {

// 	sql := `SELECT
// 				C_INTERVIEW.id,
// 				C_INTERVIEW.name,
// 				C_INTERVIEW.position,
// 				C_USER.uid,
// 				C_USER.display_name
//           	FROM
//           		C_INTERVIEW
//             LEFT JOIN
//             	C_USER ON C_INTERVIEW.uid = C_USER.uid
//           	WHERE
//           		C_INTERVIEW.published = 1
//           	AND
//           		C_INTERVIEW.id = ?`

// 	var block BasicBlock

// 	row := db.QueryRow(sql, id)
// 	err := row.Scan()
// 	if err != nil {
// 		log.Printf("%s", err)
// 	}

// 	return interview
// }
