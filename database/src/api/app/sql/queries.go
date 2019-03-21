package sql

//user_id regexp '^22+'

const (
	SelectNewUsersMLBLimited = `
SELECT  
	MIN(date_created) AS date_investing, 
	user_id AS user_id
FROM  
	user_status_history 
WHERE 
	new_user_status_id = 'investing' AND 
    date_created >= ? AND 
    date_created < ? 
    
GROUP BY  
	user_id 
ORDER BY  
	1, 2 ASC LIMIT ? OFFSET ?
`

	CountNewUsersMLB = `
    
      SELECT 
    COUNT(*)
    FROM
    (SELECT 
        MIN(date_created) AS date_investing
    FROM
        user_status_history
    WHERE
        new_user_status_id = 'investing'
            AND date_created >= ?
            AND date_created < ?
    GROUP BY user_id) AS users
      
`
)
