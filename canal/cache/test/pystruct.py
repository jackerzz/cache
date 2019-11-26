from faker import Faker

faker = Faker("zh_CN")






def run():
	print(faker.texts())





if __name__ == '__main__':
	run()