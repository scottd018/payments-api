VERSION ?= development
API_VERSION ?= v1
build:
	docker build . -t payments-api:$(VERSION)

run:
	cd test && docker-compose up

clean:
	cd test && docker-compose down -v --rmi all --remove-orphans

.PHONY: docs
docs:
	swag init \
		--dir api/$(API_VERSION) \
		--output api/$(API_VERSION)/docs \
		--generalInfo api.go \
		--parseDependencyLevel 3

#
# test harnesses
#

# payment children
test.create.child.%:
	@echo "creating $*" && curl -X POST -d @test/data/$*.json localhost:8080/api/v1/$* | jq -r && echo

test.list.child.%:
	@echo "listing $*" && curl -X GET localhost:8080/api/v1/$* | jq -r && echo

test.read.child.%:
	@echo "reading $*" && curl -X GET localhost:8080/api/v1/$*/123456 | jq -r && echo

test.update.child.%:
	@echo "updating $*" && curl -X PUT -d @test/data/$*-update.json localhost:8080/api/v1/$*/123456 | jq -r && echo

test.delete.child.%:
	@echo "deleting $*" && curl -X DELETE localhost:8080/api/v1/$*/123456 | jq -r && echo

test.child.%: test.create.child.% \
	test.list.child.% \
	test.read.child.% \
	test.update.child.% \
	test.delete.child.%
	
	@echo

# payment
test.create.payment.children: test.create.child.account_types \
	test.create.child.source_accounts \
	test.create.child.destination_accounts
	
	@echo

test.delete.payment.children: test.delete.child.source_accounts \
	test.delete.child.destination_accounts \
	test.delete.child.account_types

	@echo

test.create.payment:
	@echo "creating payment" && curl -X POST -d @test/data/payments.json localhost:8080/api/v1/payments | jq -r && echo

test.list.payment:
	@echo "listing payments" && curl -X GET localhost:8080/api/v1/payments | jq -r && echo

test.read.payment:
	@echo "reading payment" && curl -X GET localhost:8080/api/v1/payments/123456 | jq -r && echo

test.update.payment:
	@echo "updating payment" && curl -X PUT -d @test/data/payments-update.json localhost:8080/api/v1/payments/123456 | jq -r && echo

test.delete.payment:
	@echo "deleting payment" && curl -X DELETE localhost:8080/api/v1/payments/123456 | jq -r && echo

test.payment: test.create.payment.children \
	test.create.payment \
	test.list.payment \
	test.read.payment \
	test.update.payment \
	test.delete.payment \
	test.delete.payment.children
