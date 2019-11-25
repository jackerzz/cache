import pymysql
from faker import Faker


def run():
	fake = Faker("zh-CN")
	for i in range(500):
		sql = """insert into user(username,password,address)
	    values('%s','%s','%s')""" \
		      % (fake.user_name(), fake.password(special_chars=False), fake.address())
		cursor.execute(sql)


# sql = """
# create table user(
# id int PRIMARY KEY auto_increment,
# username VARCHAR(20),
# password VARCHAR(20),
# address VARCHAR(35)
# )
# """
# cursor.execute(sql)
# https://faker.readthedocs.io/en/master/locales/zh_CN.html
if __name__ == '__main__':
	conn = pymysql.connect(host="192.168.225.129", port=3306, user="root", password="root", db="go", charset="utf8")
	cursor = conn.cursor()
	
	run()
	
	
	conn.commit()
	cursor.close()
	conn.close()
