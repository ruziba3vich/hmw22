package main

/*
Uy vazifa
Github da yangi repository yarating
Local directory ga clone qilib oling
Project ichida main package dan tashqari git digan package yarating
git package ichida GetUserName digan function yarating
Shu holda yangi feature/add-username digan branch yarating va main branch dan shu branchga o'tib o'ling
GetUserName function ichida git degi username noming ni oqib oling (ishora: exec.Command method)
username ni yangi faylda yozib qoying
Barcha orgarishlarni git add hamda git commit qilib qoying
Shu holda yangi feature/add-user-emal digan branch yarating va feature/add-username dan shu branchga o'tib o'ling
GetUserEmail function yarating hamda git user emailni exec.Command orqali o'qib oling
user email ni han yaratilgan faylga append qilib qoyin
Shu branch degi barcha orgarishlarni git add hamda git commit qilib qoying
feature/add-username branchi ga qaytib oting
feature/add-user-emal degi o'zgarishlarni feature/add-username branch da merge qivoling
Natijada github repositoryni linkini tashlang

*/

import (
	"fmt"

	"github.com/ruziba3vich/hmw22/repo"
)

func main() {
	username, err := repo.GetUserEmail()
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	fmt.Println("username:", username)
}
