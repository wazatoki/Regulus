/*

maker はメーカーマスターメンテナンスのためのパッケージを提供します。

*/

package maker

import (
	makerEntity "regulus/app/domain/entities/maker"
)

/*

Create は新規作成時のユーズケースです。

正常に保存が終了した場合はMaker構造体のIDにidを代入して返却します。
保存時に例外が発生した場合はnilを返却します。
*/
func Create(makerEn makerEntity.Maker, p persistance) *makerEntity.Maker {

	var err error

	makerEn.ID, err = p.Insert(&makerEn)
	if err != nil {
		return nil
	}
	return &makerEn
}

/*

Update は更新時のユーズケースです。

正常に保存が終了した場合は引数で渡されたMaker構造体をそのまま返却します。
保存時に例外が発生した場合はnilを返却します。
*/
func Update(makerEn makerEntity.Maker, p persistance) *makerEntity.Maker {

	err := p.Update(&makerEn)
	if err != nil {
		return nil
	}

	return &makerEn
}

/*

Delete は削除時のユーズケースです。複数アイテムを削除します。

削除に失敗したMaker構造体のリストを返却します。
*/
func Delete(idList []string, p persistance) *[]makerEntity.Maker {

	errResult := []makerEntity.Maker{}

	for _, id := range idList {

		err := p.Delete(id)
		if err != nil {
			maker, err := p.SelectByID(id)
			if err == nil {
				errResult = append(errResult, *maker)
			}
		}
	}

	return &errResult
}

/*

FindByID はID検索時のユースケースです。

*/
func FindByID(id string, p persistance) *makerEntity.Maker {

	result, err := p.SelectByID(id)

	if err != nil {
		return nil
	}

	return result
}

/*

Find は検索時のユースケースです。条件指定がない場合は全件検索の結果を返します。

*/
func Find() {}
