pipeline:
	@fly sp -t local -c ci/pipeline.yml -p concourse-test
