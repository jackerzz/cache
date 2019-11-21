package dao

import (
	"cache/cache/logic/model"
	"cache/cache/logic/public/imctx"
	"cache/cache/logic/public/logger"
	"database/sql"
)

type userDao struct{}

var  UserDao  = new(userDao)
func (*userDao)Add(ctx *imctx.Context,user model.User) (int64, error) {
	result, err := DBCli.Exec("insert ignore into user(app_id,user_id,nickname,sex,avatar_url,extra) values(?,?,?,?,?,?)",
		user.AppId, user.UserId, user.Nickname, user.Sex, user.AvatarUrl, user.Extra)
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		logger.Sugar.Error(err)
		return 0, err
	}
	return affected, nil
}

func (*userDao)Get(ctx *imctx.Context,appid, userid int64) (*model.User, error){
	row := DBCli.QueryRow("select nickname,sex,avatar_url,extra,create_time,update_time from user where app_id = ? and user_id = ?",appid,userid)
	user := model.User{
		AppId: appid,
		UserId: userid,
	}
	err := row.Scan(&user.Nickname,&user.Sex,&user.AvatarUrl,&user.Extra,&user.CreateTime,&user.UpdateTime)
	if err !=nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	if err == sql.ErrNoRows{
		return nil, nil
	}
	return &user, err

}

func (*userDao)Update(ctx *imctx.Context, user model.User) error {
	_, err := DBCli.Exec("update user set nickname = ?, sex = ?, avatar_url = ?,extra = ? where app_id = and user_id = ?",
		user.Nickname, user.Sex, user.AvatarUrl, user.Extra, user.AppId, user.UserId)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}