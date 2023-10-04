package collector

/*
For now at least, we are keeping these as v1beta1 to have some element of regression testing, now that we've flipped
the "default" to v1.
*/

const prYaml = `
apiVersion: tekton.dev/v1beta1
kind: PipelineRun
metadata:
  generateName: human-resources-on-pull-request-
  annotations:
    appstudio.openshift.io/snapshot: test-application-kgccv
    pipelinesascode.tekton.dev/state: completed
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: human-resources-on-pull-request-rlrj8
  uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  creationTimestamp: '2023-08-25T14:56:26Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:labels':
            'f:tekton.dev/pipeline': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T14:56:26Z'
    - apiVersion: tekton.dev/v1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:generateName': {}
          'f:labels':
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:pipelineSpec':
            .: {}
            'f:finally': {}
            'f:params': {}
            'f:results': {}
            'f:tasks': {}
            'f:workspaces': {}
          'f:taskRunTemplate': {}
          'f:workspaces': {}
      manager: pipelines-as-code-controller
      operation: Update
      time: '2023-08-25T14:56:26Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/record': {}
            'f:results.tekton.dev/result': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T14:56:32Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          'f:childReferences': {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:pipelineResults': {}
          'f:pipelineSpec':
            .: {}
            'f:finally': {}
            'f:params': {}
            'f:results': {}
            'f:tasks': {}
            'f:workspaces': {}
          'f:skippedTasks': {}
          'f:startTime': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T15:01:09Z'
    - apiVersion: tekton.dev/v1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:pipelinesascode.tekton.dev/state': {}
          'f:finalizers':
            'v:"pipelinesascode.tekton.dev"': {}
          'f:labels':
            'f:pipelinesascode.tekton.dev/state': {}
      manager: pipelines-as-code-watcher
      operation: Update
      time: '2023-08-25T15:01:10Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:appstudio.openshift.io/snapshot': {}
      manager: manager
      operation: Update
      time: '2023-08-25T15:01:11Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev/pipelinerun"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T15:01:11Z'
  namespace: rhtapuser-tenant
  finalizers:
    - chains.tekton.dev/pipelinerun
    - pipelinesascode.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: completed
    appstudio.openshift.io/component: human-resources
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  params:
    - name: dockerfile
      value: >-
        https://raw.githubusercontent.com/devfile-samples/devfile-sample-java-springboot-basic/main/docker/Dockerfile
    - name: git-url
      value: 'https://github.com/jeff-phillips-18/human-resources'
    - name: image-expires-after
      value: 5d
    - name: output-image
      value: >-
        quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    - name: path-context
      value: .
    - name: revision
      value: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  pipelineSpec:
    finally:
      - name: show-sbom
        params:
          - name: IMAGE_URL
            value: $(tasks.build-container.results.IMAGE_URL)
        taskRef:
          kind: Task
          params:
            - name: name
              value: show-sbom
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-show-sbom:0.1@sha256:7db0af43dcebaeb33e34413148370e17078c30fd2fc78fb84c8941b444199f36
            - name: kind
              value: task
          resolver: bundles
      - name: show-summary
        params:
          - name: pipelinerun-name
            value: $(context.pipelineRun.name)
          - name: git-url
            value: >-
              $(tasks.clone-repository.results.url)?rev=$(tasks.clone-repository.results.commit)
          - name: image-url
            value: $(params.output-image)
          - name: build-task-status
            value: $(tasks.build-container.status)
        taskRef:
          kind: Task
          params:
            - name: name
              value: summary
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-summary:0.1@sha256:e69f53a3991d7088d8aa2827365ab761ab7524d4269f296b4a78b0f085789d30
            - name: kind
              value: task
          resolver: bundles
    params:
      - description: Source Repository URL
        name: git-url
        type: string
      - default: ''
        description: Revision of the Source Repository
        name: revision
        type: string
      - description: Fully Qualified Output Image
        name: output-image
        type: string
      - default: .
        description: The path to your source code
        name: path-context
        type: string
      - default: Dockerfile
        description: Path to the Dockerfile
        name: dockerfile
        type: string
      - default: 'false'
        description: Force rebuild image
        name: rebuild
        type: string
      - default: 'false'
        description: Skip checks against built image
        name: skip-checks
        type: string
      - default: 'true'
        description: 'Skip optional checks, set false if you want to run optional checks'
        name: skip-optional
        type: string
      - default: 'false'
        description: Execute the build with network isolation
        name: hermetic
        type: string
      - default: ''
        description: Build dependencies to be prefetched by Cachi2
        name: prefetch-input
        type: string
      - default: 'false'
        description: Java build
        name: java
        type: string
      - default: ''
        description: Snyk Token Secret Name
        name: snyk-secret
        type: string
      - default: ''
        description: >-
          Image tag expiration time, time values could be something like 1h, 2d,
          3w for hours, days, and weeks, respectively.
        name: image-expires-after
        type: string
    results:
      - description: ''
        name: IMAGE_URL
        value: $(tasks.build-container.results.IMAGE_URL)
      - description: ''
        name: IMAGE_DIGEST
        value: $(tasks.build-container.results.IMAGE_DIGEST)
      - description: ''
        name: CHAINS-GIT_URL
        value: $(tasks.clone-repository.results.url)
      - description: ''
        name: CHAINS-GIT_COMMIT
        value: $(tasks.clone-repository.results.commit)
      - description: ''
        name: JAVA_COMMUNITY_DEPENDENCIES
        value: $(tasks.build-container.results.JAVA_COMMUNITY_DEPENDENCIES)
    tasks:
      - name: init
        params:
          - name: image-url
            value: $(params.output-image)
          - name: rebuild
            value: $(params.rebuild)
          - name: skip-checks
            value: $(params.skip-checks)
          - name: skip-optional
            value: $(params.skip-optional)
          - name: pipelinerun-name
            value: $(context.pipelineRun.name)
          - name: pipelinerun-uid
            value: $(context.pipelineRun.uid)
        taskRef:
          kind: Task
          params:
            - name: name
              value: init
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-init:0.1@sha256:26586a7ef08c3e86dfdaf0a5cc38dd3d70c4c02db1331b469caaed0a0f5b3d86
            - name: kind
              value: task
          resolver: bundles
      - name: clone-repository
        params:
          - name: url
            value: $(params.git-url)
          - name: revision
            value: $(params.revision)
        runAfter:
          - init
        taskRef:
          kind: Task
          params:
            - name: name
              value: git-clone
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-git-clone:0.1@sha256:1f84973a21aabea38434b1f663abc4cb2d86565a9c7aae1f90decb43a8fa48eb
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(tasks.init.results.build)
            operator: in
            values:
              - 'true'
        workspaces:
          - name: output
            workspace: workspace
          - name: basic-auth
            workspace: git-auth
      - name: prefetch-dependencies
        params:
          - name: input
            value: $(params.prefetch-input)
        runAfter:
          - clone-repository
        taskRef:
          kind: Task
          params:
            - name: name
              value: prefetch-dependencies
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-prefetch-dependencies:0.1@sha256:c7b7f13d5d2a1545e95c2d56521327001d56ba54645900db41aa414607eff1e5
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(params.hermetic)
            operator: in
            values:
              - 'true'
        workspaces:
          - name: source
            workspace: workspace
      - name: build-container
        params:
          - name: IMAGE
            value: $(params.output-image)
          - name: DOCKERFILE
            value: $(params.dockerfile)
          - name: CONTEXT
            value: $(params.path-context)
          - name: HERMETIC
            value: $(params.hermetic)
          - name: PREFETCH_INPUT
            value: $(params.prefetch-input)
          - name: IMAGE_EXPIRES_AFTER
            value: $(params.image-expires-after)
          - name: COMMIT_SHA
            value: $(tasks.clone-repository.results.commit)
        runAfter:
          - prefetch-dependencies
        taskRef:
          kind: Task
          params:
            - name: name
              value: buildah
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-buildah:0.1@sha256:e607665f13adadbd4a8d0b32768fc1b24a90884d867ecb681e15c5bc25434f71
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(tasks.init.results.build)
            operator: in
            values:
              - 'true'
        workspaces:
          - name: source
            workspace: workspace
      - name: inspect-image
        params:
          - name: IMAGE_URL
            value: $(tasks.build-container.results.IMAGE_URL)
          - name: IMAGE_DIGEST
            value: $(tasks.build-container.results.IMAGE_DIGEST)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: inspect-image
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-inspect-image:0.1@sha256:bbc286f0a2ad94e671ceb9d0f1debd96f36b8c38c1147c5030957820b4125fc6
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(params.skip-checks)
            operator: in
            values:
              - 'false'
        workspaces:
          - name: source
            workspace: workspace
      - name: label-check
        runAfter:
          - inspect-image
        taskRef:
          kind: Task
          params:
            - name: name
              value: label-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-label-check:0.1@sha256:0c0739fdda24cd1e3587bbab9b07d4493efc21884baac7723f4b446e95bf1fd3
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(params.skip-checks)
            operator: in
            values:
              - 'false'
        workspaces:
          - name: workspace
            workspace: workspace
      - name: optional-label-check
        params:
          - name: POLICY_NAMESPACE
            value: optional_checks
        runAfter:
          - inspect-image
        taskRef:
          kind: Task
          params:
            - name: name
              value: label-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-label-check:0.1@sha256:0c0739fdda24cd1e3587bbab9b07d4493efc21884baac7723f4b446e95bf1fd3
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(params.skip-optional)
            operator: in
            values:
              - 'false'
        workspaces:
          - name: workspace
            workspace: workspace
      - name: deprecated-base-image-check
        params:
          - name: BASE_IMAGES_DIGESTS
            value: $(tasks.build-container.results.BASE_IMAGES_DIGESTS)
        taskRef:
          kind: Task
          params:
            - name: name
              value: deprecated-image-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-deprecated-image-check:0.2@sha256:58d16de95b4ca597f7f860fb85d6206e549910fa7a8d2a2cc229558f791ad329
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(params.skip-checks)
            operator: in
            values:
              - 'false'
        workspaces:
          - name: test-ws
            workspace: workspace
      - name: clair-scan
        params:
          - name: image-digest
            value: $(tasks.build-container.results.IMAGE_DIGEST)
          - name: image-url
            value: $(tasks.build-container.results.IMAGE_URL)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: clair-scan
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-clair-scan:0.1@sha256:c5602d9d6dd797da98e98fde8471ea55a788c30f74f2192807910ce5436e9b66
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(params.skip-checks)
            operator: in
            values:
              - 'false'
      - name: sast-snyk-check
        params:
          - name: SNYK_SECRET
            value: $(params.snyk-secret)
        runAfter:
          - clone-repository
        taskRef:
          kind: Task
          params:
            - name: name
              value: sast-snyk-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-sast-snyk-check:0.1@sha256:9dcd450b454705b9fe22c5f8f7bb7305cebc3cb73e783b85e047f7e721994189
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(params.skip-checks)
            operator: in
            values:
              - 'false'
          - input: $(params.snyk-secret)
            operator: notin
            values:
              - ''
        workspaces:
          - name: workspace
            workspace: workspace
      - name: clamav-scan
        params:
          - name: image-digest
            value: $(tasks.build-container.results.IMAGE_DIGEST)
          - name: image-url
            value: $(tasks.build-container.results.IMAGE_URL)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: clamav-scan
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-clamav-scan:0.1@sha256:cd4e301dd849cbdf7b8e38fd8f4915970b5b60174770df632a6b38ea93028d44
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(params.skip-checks)
            operator: in
            values:
              - 'false'
      - name: sbom-json-check
        params:
          - name: IMAGE_URL
            value: $(tasks.build-container.results.IMAGE_URL)
          - name: IMAGE_DIGEST
            value: $(tasks.build-container.results.IMAGE_DIGEST)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: sbom-json-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-sbom-json-check:0.1@sha256:397cb2fb20f413dec9653134231bec86edb80806a3441081fbf473677fc40917
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(params.skip-checks)
            operator: in
            values:
              - 'false'
    workspaces:
      - name: workspace
      - name: git-auth
        optional: true
  serviceAccountName: appstudio-pipeline
  timeouts:
    pipeline: 1h0m0s
  workspaces:
    - name: workspace
      volumeClaimTemplate:
        metadata:
          creationTimestamp: null
        spec:
          accessModes:
            - ReadWriteOnce
          resources:
            requests:
              storage: 1Gi
        status: {}
    - name: git-auth
      secret:
        secretName: pac-gitauth-pwvj
status:
  childReferences:
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: human-resources-on-pull-request-rlrj8-init
      pipelineTaskName: init
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: human-resources-on-pull-request-rlrj8-clone-repository
      pipelineTaskName: clone-repository
      whenExpressions:
        - input: $(tasks.init.results.build)
          operator: in
          values:
            - 'true'
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: human-resources-on-pull-request-rlrj8-build-container
      pipelineTaskName: build-container
      whenExpressions:
        - input: $(tasks.init.results.build)
          operator: in
          values:
            - 'true'
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: human-resources-on-pull-request-rlrj8-inspect-image
      pipelineTaskName: inspect-image
      whenExpressions:
        - input: 'false'
          operator: in
          values:
            - 'false'
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: human-resources-on-pull-request-rlrj8-label-check
      pipelineTaskName: label-check
      whenExpressions:
        - input: 'false'
          operator: in
          values:
            - 'false'
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: humeaa4363f279f84d4e98a2fee1964cf0b-deprecated-base-image-check
      pipelineTaskName: deprecated-base-image-check
      whenExpressions:
        - input: 'false'
          operator: in
          values:
            - 'false'
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: human-resources-on-pull-request-rlrj8-clair-scan
      pipelineTaskName: clair-scan
      whenExpressions:
        - input: 'false'
          operator: in
          values:
            - 'false'
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: human-resources-on-pull-request-rlrj8-clamav-scan
      pipelineTaskName: clamav-scan
      whenExpressions:
        - input: 'false'
          operator: in
          values:
            - 'false'
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: human-resources-on-pull-request-rlrj8-sbom-json-check
      pipelineTaskName: sbom-json-check
      whenExpressions:
        - input: 'false'
          operator: in
          values:
            - 'false'
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: human-resources-on-pull-request-rlrj8-show-sbom
      pipelineTaskName: show-sbom
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: human-resources-on-pull-request-rlrj8-show-summary
      pipelineTaskName: show-summary
  completionTime: '2023-08-25T15:01:09Z'
  conditions:
    - lastTransitionTime: '2023-08-25T15:01:09Z'
      message: 'Tasks Completed: 11 (Failed: 0, Cancelled 0), Skipped: 3'
      reason: Completed
      status: 'True'
      type: Succeeded
  pipelineResults:
    - name: IMAGE_URL
      value: >-
        quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    - name: IMAGE_DIGEST
      value: 'sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d'
    - name: CHAINS-GIT_URL
      value: 'https://github.com/jeff-phillips-18/human-resources'
    - name: CHAINS-GIT_COMMIT
      value: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    - name: JAVA_COMMUNITY_DEPENDENCIES
      value: ''
  pipelineSpec:
    finally:
      - name: show-sbom
        params:
          - name: IMAGE_URL
            value: $(tasks.build-container.results.IMAGE_URL)
        taskRef:
          kind: Task
          params:
            - name: name
              value: show-sbom
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-show-sbom:0.1@sha256:7db0af43dcebaeb33e34413148370e17078c30fd2fc78fb84c8941b444199f36
            - name: kind
              value: task
          resolver: bundles
        status:
          completionTime: '2023-08-25T15:01:09Z'
          conditions:
            - lastTransitionTime: '2023-08-25T15:01:09Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: human-resources-on-pull-request-rlrj8-show-sbom-pod
          startTime: '2023-08-25T15:01:02Z'
          steps:
            - container: step-show-sbom
              imageID: >-
                quay.io/redhat-appstudio/cosign@sha256:c883d6f8d39148f2cea71bff4622d196d89df3e510f36c140c097b932f0dd5d5
              name: show-sbom
              terminated:
                containerID: >-
                  cri-o://6c057c683644abb99da528594de18b2f7d56244c44d04e1c92f6eba1e0e083e7
                exitCode: 0
                finishedAt: '2023-08-25T15:01:08Z'
                reason: Completed
                startedAt: '2023-08-25T15:01:07Z'
          taskSpec:
            description: >-
              Shows the Software Bill of Materials (SBOM) generated for the
              built image in CyloneDX JSON format.
            params:
              - description: Fully qualified image name to show SBOM for.
                name: IMAGE_URL
                type: string
            steps:
              - env:
                  - name: IMAGE_URL
                    value: >-
                      quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                image: 'quay.io/redhat-appstudio/cosign:v2.1.1'
                name: show-sbom
                resources: {}
                script: |
                  #!/busybox/sh
                  cosign download sbom $IMAGE_URL 2>err
                  RET=$?
                  if [ $RET -ne 0 ]; then
                    echo Failed to get SBOM >&2
                    cat err >&2
                  fi
                  exit $RET
          duration: 7s
          reason: Succeeded
      - name: show-summary
        params:
          - name: pipelinerun-name
            value: human-resources-on-pull-request-rlrj8
          - name: git-url
            value: >-
              $(tasks.clone-repository.results.url)?rev=$(tasks.clone-repository.results.commit)
          - name: image-url
            value: >-
              quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          - name: build-task-status
            value: $(tasks.build-container.status)
        taskRef:
          kind: Task
          params:
            - name: name
              value: summary
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-summary:0.1@sha256:e69f53a3991d7088d8aa2827365ab761ab7524d4269f296b4a78b0f085789d30
            - name: kind
              value: task
          resolver: bundles
        status:
          completionTime: '2023-08-25T15:01:08Z'
          conditions:
            - lastTransitionTime: '2023-08-25T15:01:08Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: human-resources-on-pull-request-rlrj8-show-summary-pod
          startTime: '2023-08-25T15:01:02Z'
          steps:
            - container: step-appstudio-summary
              imageID: >-
                registry.access.redhat.com/ubi9/ubi-minimal@sha256:dc02c6aa8199beb8ed13312d7116a94aa87b5412886bbe358845d3f0626c0f1e
              name: appstudio-summary
              terminated:
                containerID: >-
                  cri-o://aa6b8e16a118ea8f712a46a380d40696d5975f7f73472f0d5276f45a20c311dc
                exitCode: 0
                finishedAt: '2023-08-25T15:01:07Z'
                reason: Completed
                startedAt: '2023-08-25T15:01:07Z'
          taskSpec:
            description: >-
              Summary Pipeline Task. Prints PipelineRun information, removes
              image repository secret used by the PipelineRun.
            params:
              - description: pipeline-run to annotate
                name: pipelinerun-name
                type: string
              - description: Git URL
                name: git-url
                type: string
              - description: Image URL
                name: image-url
                type: string
              - default: Succeeded
                description: State of build task in pipelineRun
                name: build-task-status
                type: string
            steps:
              - env:
                  - name: GIT_URL
                    value: >-
                      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                  - name: IMAGE_URL
                    value: >-
                      quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                  - name: PIPELINERUN_NAME
                    value: human-resources-on-pull-request-rlrj8
                  - name: BUILD_TASK_STATUS
                    value: Succeeded
                image: >-
                  registry.access.redhat.com/ubi9/ubi-minimal:9.2-717@sha256:dc02c6aa8199beb8ed13312d7116a94aa87b5412886bbe358845d3f0626c0f1e
                name: appstudio-summary
                resources: {}
                script: |
                  #!/usr/bin/env bash
                  echo
                  echo "Build Summary:"
                  echo
                  echo "Build repository: $GIT_URL"
                  if [ "$BUILD_TASK_STATUS" == "Succeeded" ]; then
                    echo "Generated Image is in : $IMAGE_URL"
                  fi
                  echo
                  echo End Summary
          duration: 6s
          reason: Succeeded
    params:
      - description: Source Repository URL
        name: git-url
        type: string
      - default: ''
        description: Revision of the Source Repository
        name: revision
        type: string
      - description: Fully Qualified Output Image
        name: output-image
        type: string
      - default: .
        description: The path to your source code
        name: path-context
        type: string
      - default: Dockerfile
        description: Path to the Dockerfile
        name: dockerfile
        type: string
      - default: 'false'
        description: Force rebuild image
        name: rebuild
        type: string
      - default: 'false'
        description: Skip checks against built image
        name: skip-checks
        type: string
      - default: 'true'
        description: 'Skip optional checks, set false if you want to run optional checks'
        name: skip-optional
        type: string
      - default: 'false'
        description: Execute the build with network isolation
        name: hermetic
        type: string
      - default: ''
        description: Build dependencies to be prefetched by Cachi2
        name: prefetch-input
        type: string
      - default: 'false'
        description: Java build
        name: java
        type: string
      - default: ''
        description: Snyk Token Secret Name
        name: snyk-secret
        type: string
      - default: ''
        description: >-
          Image tag expiration time, time values could be something like 1h, 2d,
          3w for hours, days, and weeks, respectively.
        name: image-expires-after
        type: string
    results:
      - description: ''
        name: IMAGE_URL
        value: $(tasks.build-container.results.IMAGE_URL)
      - description: ''
        name: IMAGE_DIGEST
        value: $(tasks.build-container.results.IMAGE_DIGEST)
      - description: ''
        name: CHAINS-GIT_URL
        value: $(tasks.clone-repository.results.url)
      - description: ''
        name: CHAINS-GIT_COMMIT
        value: $(tasks.clone-repository.results.commit)
      - description: ''
        name: JAVA_COMMUNITY_DEPENDENCIES
        value: $(tasks.build-container.results.JAVA_COMMUNITY_DEPENDENCIES)
    tasks:
      - name: init
        params:
          - name: image-url
            value: >-
              quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          - name: rebuild
            value: 'false'
          - name: skip-checks
            value: 'false'
          - name: skip-optional
            value: 'true'
          - name: pipelinerun-name
            value: human-resources-on-pull-request-rlrj8
          - name: pipelinerun-uid
            value: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
        taskRef:
          kind: Task
          params:
            - name: name
              value: init
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-init:0.1@sha256:26586a7ef08c3e86dfdaf0a5cc38dd3d70c4c02db1331b469caaed0a0f5b3d86
            - name: kind
              value: task
          resolver: bundles
        status:
          completionTime: '2023-08-25T14:56:37Z'
          conditions:
            - lastTransitionTime: '2023-08-25T14:56:37Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: human-resources-on-pull-request-rlrj8-init-pod
          startTime: '2023-08-25T14:56:31Z'
          steps:
            - container: step-init
              imageID: >-
                registry.redhat.io/openshift4/ose-cli@sha256:73df37794ffff7de1101016c23dc623e4990810390ebdabcbbfa065214352c7c
              name: init
              terminated:
                containerID: >-
                  cri-o://c75af9b7fb35fad7947213d9151d0922396e426f3c8d43f92689c8349b8d1e04
                exitCode: 0
                finishedAt: '2023-08-25T14:56:37Z'
                message: >-
                  [{"key":"build","value":"true","type":1},{"key":"container-registry-secret","value":"unused\n","type":1}]
                reason: Completed
                startedAt: '2023-08-25T14:56:36Z'
          taskResults:
            - name: build
              type: string
              value: 'true'
            - name: container-registry-secret
              type: string
              value: |
                unused
          taskSpec:
            description: >-
              Initialize Pipeline Task, include flags for rebuild and auth.
              Generates image repository secret used by the PipelineRun.
            params:
              - description: Image URL for build by PipelineRun
                name: image-url
                type: string
              - default: 'false'
                description: Rebuild the image if exists
                name: rebuild
                type: string
              - default: 'false'
                description: Skip checks against built image
                name: skip-checks
                type: string
              - default: 'true'
                description: >-
                  Skip optional checks, set false if you want to run optional
                  checks
                name: skip-optional
                type: string
              - default: ''
                description: 'unused, should be removed in next task version'
                name: pipelinerun-name
                type: string
              - default: ''
                description: 'unused, should be removed in next task version'
                name: pipelinerun-uid
                type: string
            results:
              - description: Defines if the image in param image-url should be built
                name: build
                type: string
              - description: 'unused, should be removed in next task version'
                name: container-registry-secret
                type: string
            steps:
              - env:
                  - name: IMAGE_URL
                    value: >-
                      quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                  - name: REBUILD
                    value: 'false'
                  - name: SKIP_CHECKS
                    value: 'false'
                  - name: SKIP_OPTIONAL
                    value: 'true'
                image: >-
                  registry.redhat.io/openshift4/ose-cli:4.13@sha256:73df37794ffff7de1101016c23dc623e4990810390ebdabcbbfa065214352c7c
                name: init
                resources: {}
                script: >
                  #!/bin/bash

                  echo "Build Initialize: $IMAGE_URL"

                  echo


                  echo "Determine if Image Already Exists"

                  # Build the image when image does not exists or rebuild is set
                  to true

                  if ! oc image info $IMAGE_URL &>/dev/null || [ "$REBUILD" ==
                  "true" ] || [ "$SKIP_CHECKS" == "false" ]; then
                    echo -n "true" > /tekton/results/build
                  else
                    echo -n "false" > /tekton/results/build
                  fi

                  echo unused > /tekton/results/container-registry-secret
          duration: 6s
          reason: Succeeded
      - name: clone-repository
        params:
          - name: url
            value: 'https://github.com/jeff-phillips-18/human-resources'
          - name: revision
            value: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
        runAfter:
          - init
        taskRef:
          kind: Task
          params:
            - name: name
              value: git-clone
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-git-clone:0.1@sha256:1f84973a21aabea38434b1f663abc4cb2d86565a9c7aae1f90decb43a8fa48eb
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(tasks.init.results.build)
            operator: in
            values:
              - 'true'
        workspaces:
          - name: output
            workspace: workspace
          - name: basic-auth
            workspace: git-auth
        status:
          completionTime: '2023-08-25T14:56:54Z'
          conditions:
            - lastTransitionTime: '2023-08-25T14:56:54Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: human-resources-on-pull-request-rlrj8-clone-repository-pod
          startTime: '2023-08-25T14:56:37Z'
          steps:
            - container: step-clone
              imageID: >-
                registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8@sha256:2fa0b06d52b04f377c696412e19307a9eff27383f81d87aae0b4f71672a1cd0b
              name: clone
              terminated:
                containerID: >-
                  cri-o://ad6aebe8612f8776692ced04d7726d804642f3a5319004d9cd7e2b1bc780ed4b
                exitCode: 0
                finishedAt: '2023-08-25T14:56:54Z'
                message: >-
                  [{"key":"commit","value":"b9651320a8c6fdd24cf35d2a1f8ea0a245235442","type":1},{"key":"url","value":"https://github.com/jeff-phillips-18/human-resources","type":1}]
                reason: Completed
                startedAt: '2023-08-25T14:56:54Z'
            - container: step-symlink-check
              imageID: >-
                registry.redhat.io/ubi9@sha256:089bd3b82a78ac45c0eed231bb58bfb43bfcd0560d9bba240fc6355502c92976
              name: symlink-check
              terminated:
                containerID: >-
                  cri-o://e6a39307c093b7d07b08bcc1204e13b30e89af83a87bce7f18283ef283b078d0
                exitCode: 0
                finishedAt: '2023-08-25T14:56:54Z'
                message: >-
                  [{"key":"commit","value":"b9651320a8c6fdd24cf35d2a1f8ea0a245235442","type":1},{"key":"url","value":"https://github.com/jeff-phillips-18/human-resources","type":1}]
                reason: Completed
                startedAt: '2023-08-25T14:56:54Z'
          taskResults:
            - name: commit
              type: string
              value: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
            - name: url
              type: string
              value: 'https://github.com/jeff-phillips-18/human-resources'
          taskSpec:
            description: >-
              The git-clone Task will clone a repo from the provided url into
              the output Workspace. By default the repo will be cloned into the
              root of your Workspace.
            params:
              - description: Repository URL to clone from.
                name: url
                type: string
              - default: ''
                description: 'Revision to checkout. (branch, tag, sha, ref, etc...)'
                name: revision
                type: string
              - default: ''
                description: Refspec to fetch before checking out revision.
                name: refspec
                type: string
              - default: 'true'
                description: Initialize and fetch git submodules.
                name: submodules
                type: string
              - default: '1'
                description: >-
                  Perform a shallow clone, fetching only the most recent N
                  commits.
                name: depth
                type: string
              - default: 'true'
                description: >-
                  Set the http.sslVerify global git config. Setting this to
                  false is not advised unless you are sure that you trust your
                  git remote.
                name: sslVerify
                type: string
              - default: ''
                description: >-
                  Subdirectory inside the output Workspace to clone the repo
                  into.
                name: subdirectory
                type: string
              - default: ''
                description: >-
                  Define the directory patterns to match or exclude when
                  performing a sparse checkout.
                name: sparseCheckoutDirectories
                type: string
              - default: 'true'
                description: >-
                  Clean out the contents of the destination directory if it
                  already exists before cloning.
                name: deleteExisting
                type: string
              - default: ''
                description: HTTP proxy server for non-SSL requests.
                name: httpProxy
                type: string
              - default: ''
                description: HTTPS proxy server for SSL requests.
                name: httpsProxy
                type: string
              - default: ''
                description: Opt out of proxying HTTP/HTTPS requests.
                name: noProxy
                type: string
              - default: 'true'
                description: >-
                  Log the commands that are executed during git-clone's
                  operation.
                name: verbose
                type: string
              - default: >-
                  registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8:v1.8.2-8@sha256:a538c423e7a11aae6ae582a411fdb090936458075f99af4ce5add038bb6983e8
                description: The image providing the git-init binary that this Task runs.
                name: gitInitImage
                type: string
              - default: /tekton/home
                description: >
                  Absolute path to the user's home directory. Set this
                  explicitly if you are running the image as a non-root user or
                  have overridden

                  the gitInitImage param with an image containing custom user
                  configuration.
                name: userHome
                type: string
              - default: 'true'
                description: >
                  Check symlinks in the repo. If they're pointing outside of the
                  repo, the build will fail.
                name: enableSymlinkCheck
                type: string
            results:
              - description: The precise commit SHA that was fetched by this Task.
                name: commit
                type: string
              - description: The precise URL that was fetched by this Task.
                name: url
                type: string
            steps:
              - env:
                  - name: HOME
                    value: /tekton/home
                  - name: PARAM_URL
                    value: 'https://github.com/jeff-phillips-18/human-resources'
                  - name: PARAM_REVISION
                    value: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                  - name: PARAM_REFSPEC
                  - name: PARAM_SUBMODULES
                    value: 'true'
                  - name: PARAM_DEPTH
                    value: '1'
                  - name: PARAM_SSL_VERIFY
                    value: 'true'
                  - name: PARAM_SUBDIRECTORY
                  - name: PARAM_DELETE_EXISTING
                    value: 'true'
                  - name: PARAM_HTTP_PROXY
                  - name: PARAM_HTTPS_PROXY
                  - name: PARAM_NO_PROXY
                  - name: PARAM_VERBOSE
                    value: 'true'
                  - name: PARAM_SPARSE_CHECKOUT_DIRECTORIES
                  - name: PARAM_USER_HOME
                    value: /tekton/home
                  - name: WORKSPACE_OUTPUT_PATH
                    value: /workspace/output
                  - name: WORKSPACE_SSH_DIRECTORY_BOUND
                    value: 'false'
                  - name: WORKSPACE_SSH_DIRECTORY_PATH
                  - name: WORKSPACE_BASIC_AUTH_DIRECTORY_BOUND
                    value: 'true'
                  - name: WORKSPACE_BASIC_AUTH_DIRECTORY_PATH
                    value: /workspace/basic-auth
                image: >-
                  registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8:v1.8.2-8@sha256:a538c423e7a11aae6ae582a411fdb090936458075f99af4ce5add038bb6983e8
                name: clone
                resources: {}
                script: >
                  #!/usr/bin/env sh

                  set -eu


                  if [ "${PARAM_VERBOSE}" = "true" ] ; then
                    set -x
                  fi


                  if [ "${WORKSPACE_BASIC_AUTH_DIRECTORY_BOUND}" = "true" ] ;
                  then
                    if [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.git-credentials" ] && [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.gitconfig" ]; then
                      cp "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.git-credentials" "${PARAM_USER_HOME}/.git-credentials"
                      cp "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.gitconfig" "${PARAM_USER_HOME}/.gitconfig"
                    # Compatibility with kubernetes.io/basic-auth secrets
                    elif [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/username" ] && [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/password" ]; then
                      HOSTNAME=$(echo $PARAM_URL | awk -F/ '{print $3}')
                      echo "https://$(cat ${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/username):$(cat ${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/password)@$HOSTNAME" > "${PARAM_USER_HOME}/.git-credentials"
                      echo -e "[credential \"https://$HOSTNAME\"]\n  helper = store" > "${PARAM_USER_HOME}/.gitconfig"
                    else
                      echo "Unknown basic-auth workspace format"
                      exit 1
                    fi
                    chmod 400 "${PARAM_USER_HOME}/.git-credentials"
                    chmod 400 "${PARAM_USER_HOME}/.gitconfig"
                  fi


                  if [ "${WORKSPACE_SSH_DIRECTORY_BOUND}" = "true" ] ; then
                    cp -R "${WORKSPACE_SSH_DIRECTORY_PATH}" "${PARAM_USER_HOME}"/.ssh
                    chmod 700 "${PARAM_USER_HOME}"/.ssh
                    chmod -R 400 "${PARAM_USER_HOME}"/.ssh/*
                  fi


                  CHECKOUT_DIR="${WORKSPACE_OUTPUT_PATH}/${PARAM_SUBDIRECTORY}"


                  cleandir() {
                    # Delete any existing contents of the repo directory if it exists.
                    #
                    # We don't just "rm -rf ${CHECKOUT_DIR}" because ${CHECKOUT_DIR} might be "/"
                    # or the root of a mounted volume.
                    if [ -d "${CHECKOUT_DIR}" ] ; then
                      # Delete non-hidden files and directories
                      rm -rf "${CHECKOUT_DIR:?}"/*
                      # Delete files and directories starting with . but excluding ..
                      rm -rf "${CHECKOUT_DIR}"/.[!.]*
                      # Delete files and directories starting with .. plus any other character
                      rm -rf "${CHECKOUT_DIR}"/..?*
                    fi
                  }


                  if [ "${PARAM_DELETE_EXISTING}" = "true" ] ; then
                    cleandir
                  fi


                  test -z "${PARAM_HTTP_PROXY}" || export
                  HTTP_PROXY="${PARAM_HTTP_PROXY}"

                  test -z "${PARAM_HTTPS_PROXY}" || export
                  HTTPS_PROXY="${PARAM_HTTPS_PROXY}"

                  test -z "${PARAM_NO_PROXY}" || export
                  NO_PROXY="${PARAM_NO_PROXY}"


                  /ko-app/git-init \
                    -url="${PARAM_URL}" \
                    -revision="${PARAM_REVISION}" \
                    -refspec="${PARAM_REFSPEC}" \
                    -path="${CHECKOUT_DIR}" \
                    -sslVerify="${PARAM_SSL_VERIFY}" \
                    -submodules="${PARAM_SUBMODULES}" \
                    -depth="${PARAM_DEPTH}" \
                    -sparseCheckoutDirectories="${PARAM_SPARSE_CHECKOUT_DIRECTORIES}"
                  cd "${CHECKOUT_DIR}"

                  RESULT_SHA="$(git rev-parse HEAD)"

                  EXIT_CODE="$?"

                  if [ "${EXIT_CODE}" != 0 ] ; then
                    exit "${EXIT_CODE}"
                  fi

                  printf "%s" "${RESULT_SHA}" > "/tekton/results/commit"

                  printf "%s" "${PARAM_URL}" > "/tekton/results/url"
                securityContext:
                  runAsUser: 0
              - env:
                  - name: PARAM_ENABLE_SYMLINK_CHECK
                    value: 'true'
                  - name: PARAM_SUBDIRECTORY
                  - name: WORKSPACE_OUTPUT_PATH
                    value: /workspace/output
                image: 'registry.redhat.io/ubi9:9.2-696'
                name: symlink-check
                resources: {}
                script: |
                  #!/usr/bin/env bash
                  set -euo pipefail

                  CHECKOUT_DIR="${WORKSPACE_OUTPUT_PATH}/${PARAM_SUBDIRECTORY}"
                  check_symlinks() {
                    FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO=false
                    while read symlink
                    do
                      target=$(readlink -f "$symlink")
                      if ! [[ "$target" =~ ^$CHECKOUT_DIR ]]; then
                        echo "The cloned repository contains symlink pointing outside of the cloned repository: $symlink"
                        FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO=true
                      fi
                    done < <(find $CHECKOUT_DIR -type l -print)
                    if [ "$FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO" = true ] ; then
                      return 1
                    fi
                  }

                  if [ "${PARAM_ENABLE_SYMLINK_CHECK}" = "true" ] ; then
                    echo "Running symlink check"
                    check_symlinks
                  fi
            workspaces:
              - description: >-
                  The git repo will be cloned onto the volume backing this
                  Workspace.
                name: output
              - description: >
                  A .ssh directory with private key, known_hosts, config, etc.
                  Copied to

                  the user's home before git commands are executed. Used to
                  authenticate

                  with the git remote when performing the clone. Binding a
                  Secret to this

                  Workspace is strongly recommended over other volume types.
                name: ssh-directory
                optional: true
              - description: >
                  A Workspace containing a .gitconfig and .git-credentials file
                  or username and password.

                  These will be copied to the user's home before any git
                  commands are run. Any

                  other files in this Workspace are ignored. It is strongly
                  recommended

                  to use ssh-directory over basic-auth whenever possible and to
                  bind a

                  Secret to this Workspace over other volume types.
                name: basic-auth
                optional: true
          duration: 17s
          reason: Succeeded
      - name: prefetch-dependencies
        params:
          - name: input
            value: ''
        runAfter:
          - clone-repository
        taskRef:
          kind: Task
          params:
            - name: name
              value: prefetch-dependencies
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-prefetch-dependencies:0.1@sha256:c7b7f13d5d2a1545e95c2d56521327001d56ba54645900db41aa414607eff1e5
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'false'
            operator: in
            values:
              - 'true'
        workspaces:
          - name: source
            workspace: workspace
        status:
          reason: Skipped
      - name: build-container
        params:
          - name: IMAGE
            value: >-
              quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          - name: DOCKERFILE
            value: >-
              https://raw.githubusercontent.com/devfile-samples/devfile-sample-java-springboot-basic/main/docker/Dockerfile
          - name: CONTEXT
            value: .
          - name: HERMETIC
            value: 'false'
          - name: PREFETCH_INPUT
            value: ''
          - name: IMAGE_EXPIRES_AFTER
            value: 5d
          - name: COMMIT_SHA
            value: $(tasks.clone-repository.results.commit)
        runAfter:
          - prefetch-dependencies
        taskRef:
          kind: Task
          params:
            - name: name
              value: buildah
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-buildah:0.1@sha256:e607665f13adadbd4a8d0b32768fc1b24a90884d867ecb681e15c5bc25434f71
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(tasks.init.results.build)
            operator: in
            values:
              - 'true'
        workspaces:
          - name: source
            workspace: workspace
        status:
          completionTime: '2023-08-25T14:58:25Z'
          conditions:
            - lastTransitionTime: '2023-08-25T14:58:25Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: human-resources-on-pull-request-rlrj8-build-container-pod
          startTime: '2023-08-25T14:56:55Z'
          steps:
            - container: step-build
              imageID: >-
                quay.io/redhat-appstudio/buildah@sha256:381e9bfedd59701477621da93892106873a6951b196105d3d2d85c3f6d7b569b
              name: build
              terminated:
                containerID: >-
                  cri-o://2bee75648c19f53bdc6caa852fec6d8516391ac78dec27cd5edeb0b1355be656
                exitCode: 0
                finishedAt: '2023-08-25T14:58:01Z'
                reason: Completed
                startedAt: '2023-08-25T14:57:04Z'
            - container: step-sbom-syft-generate
              imageID: >-
                quay.io/redhat-appstudio/syft@sha256:244a17ce220a0b7a54c862c4fe3b72ce92799910c5eff8e94ac2f121fa5b4a53
              name: sbom-syft-generate
              terminated:
                containerID: >-
                  cri-o://e4824b7c6dcd0a7aee55938e60e9e2085bbd4bbd3a6442734b4855ff6a9ec919
                exitCode: 0
                finishedAt: '2023-08-25T14:58:08Z'
                reason: Completed
                startedAt: '2023-08-25T14:58:02Z'
            - container: step-analyse-dependencies-java-sbom
              imageID: >-
                quay.io/redhat-appstudio/hacbs-jvm-build-request-processor@sha256:b198cf4b33dab59ce8ac25afd4e1001390db29ca2dec83dc8a1e21b0359ce743
              name: analyse-dependencies-java-sbom
              terminated:
                containerID: >-
                  cri-o://40eac4840ea9d84672a078c389c6e872464d62614399013508708f79e35f8488
                exitCode: 0
                finishedAt: '2023-08-25T14:58:08Z'
                message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
                reason: Completed
                startedAt: '2023-08-25T14:58:08Z'
            - container: step-merge-syft-sboms
              imageID: >-
                registry.access.redhat.com/ubi9/python-39@sha256:562b4f9cd4e5abeba54c56fd2290096a71f6e10aa8c1c18f43c9d6962c2d4d41
              name: merge-syft-sboms
              terminated:
                containerID: >-
                  cri-o://78e33feea2bccd7edc2a08fc6fbcae9bf3a3d6e6477aadd7b5b792b5df1c071d
                exitCode: 0
                finishedAt: '2023-08-25T14:58:09Z'
                message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
                reason: Completed
                startedAt: '2023-08-25T14:58:09Z'
            - container: step-merge-cachi2-sbom
              imageID: >-
                quay.io/redhat-appstudio/cachi2@sha256:46097f22b57e4d48a3fce96d931e08ccfe3a3e6421362d5f9353961279078eef
              name: merge-cachi2-sbom
              terminated:
                containerID: >-
                  cri-o://3113046a374d117b72418467921a2f2a6026736a3f34d2c730c056c2bcd278f7
                exitCode: 0
                finishedAt: '2023-08-25T14:58:09Z'
                message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
                reason: Completed
                startedAt: '2023-08-25T14:58:09Z'
            - container: step-create-purl-sbom
              imageID: >-
                registry.access.redhat.com/ubi9/python-39@sha256:562b4f9cd4e5abeba54c56fd2290096a71f6e10aa8c1c18f43c9d6962c2d4d41
              name: create-purl-sbom
              terminated:
                containerID: >-
                  cri-o://5c4cd3e48bf9a27abbeee0132e6c03fd33ca72239f52f08f00601b795b6e1c4e
                exitCode: 0
                finishedAt: '2023-08-25T14:58:09Z'
                message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
                reason: Completed
                startedAt: '2023-08-25T14:58:09Z'
            - container: step-inject-sbom-and-push
              imageID: >-
                registry.access.redhat.com/ubi9/buildah@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
              name: inject-sbom-and-push
              terminated:
                containerID: >-
                  cri-o://4521e19967b852814ac76a5bd3a261dcbf1046cbdcb05091dca0aa5cc6031938
                exitCode: 0
                finishedAt: '2023-08-25T14:58:22Z'
                message: >-
                  [{"key":"BASE_IMAGES_DIGESTS","value":"registry.access.redhat.com/ubi8/openjdk-17:1.15-1.1682053058@sha256:b00f687d913b8d1e027f7eabd6765de6c8d469629bef9550f10dbf207af24fe5\nregistry.access.redhat.com/ubi8/openjdk-17-runtime:1.15-1.1682053056@sha256:f921cf1f9147e4b306908f3bcb61dd215b4a51970f8db560ede02ee6a492fa99\n","type":1},{"key":"IMAGE_DIGEST","value":"sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d","type":1},{"key":"IMAGE_URL","value":"quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442","type":1},{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]
                reason: Completed
                startedAt: '2023-08-25T14:58:10Z'
            - container: step-upload-sbom
              imageID: >-
                quay.io/redhat-appstudio/cosign@sha256:c883d6f8d39148f2cea71bff4622d196d89df3e510f36c140c097b932f0dd5d5
              name: upload-sbom
              terminated:
                containerID: >-
                  cri-o://bfd0cce431db96e5dd1fd3bbc69f777f20b64f72d05859054890fd2e992c6bd4
                exitCode: 0
                finishedAt: '2023-08-25T14:58:24Z'
                message: >-
                  [{"key":"BASE_IMAGES_DIGESTS","value":"registry.access.redhat.com/ubi8/openjdk-17:1.15-1.1682053058@sha256:b00f687d913b8d1e027f7eabd6765de6c8d469629bef9550f10dbf207af24fe5\nregistry.access.redhat.com/ubi8/openjdk-17-runtime:1.15-1.1682053056@sha256:f921cf1f9147e4b306908f3bcb61dd215b4a51970f8db560ede02ee6a492fa99\n","type":1},{"key":"IMAGE_DIGEST","value":"sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d","type":1},{"key":"IMAGE_URL","value":"quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442","type":1},{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]
                reason: Completed
                startedAt: '2023-08-25T14:58:23Z'
          taskResults:
            - name: JAVA_COMMUNITY_DEPENDENCIES
              type: string
              value: ''
            - name: BASE_IMAGES_DIGESTS
              type: string
              value: >
                registry.access.redhat.com/ubi8/openjdk-17:1.15-1.1682053058@sha256:b00f687d913b8d1e027f7eabd6765de6c8d469629bef9550f10dbf207af24fe5

                registry.access.redhat.com/ubi8/openjdk-17-runtime:1.15-1.1682053056@sha256:f921cf1f9147e4b306908f3bcb61dd215b4a51970f8db560ede02ee6a492fa99
            - name: IMAGE_DIGEST
              type: string
              value: >-
                sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d
            - name: IMAGE_URL
              type: string
              value: >-
                quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          taskSpec:
            description: >-
              Buildah task builds source code into a container image and pushes
              the image into container registry using buildah tool.

              In addition it generates a SBOM file, injects the SBOM file into
              final container image and pushes the SBOM file as separate image
              using cosign tool.

              When [Java dependency
              rebuild](https://redhat-appstudio.github.io/docs.stonesoup.io/Documentation/main/cli/proc_enabled_java_dependencies.html)
              is enabled it triggers rebuilds of Java artifacts.

              When prefetch-dependencies task was activated it is using its
              artifacts to run build in hermetic environment.
            params:
              - description: Reference of the image buildah will produce.
                name: IMAGE
                type: string
              - default: >-
                  registry.access.redhat.com/ubi9/buildah:9.0.0-19@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
                description: The location of the buildah builder image.
                name: BUILDER_IMAGE
                type: string
              - default: ./Dockerfile
                description: Path to the Dockerfile to build.
                name: DOCKERFILE
                type: string
              - default: .
                description: Path to the directory to use as context.
                name: CONTEXT
                type: string
              - default: 'true'
                description: >-
                  Verify the TLS on the registry endpoint (for push/pull to a
                  non-TLS registry)
                name: TLSVERIFY
                type: string
              - default: ''
                description: 'unused, should be removed in next task version'
                name: DOCKER_AUTH
                type: string
              - default: 'false'
                description: Determines if build will be executed without network access.
                name: HERMETIC
                type: string
              - default: ''
                description: >-
                  In case it is not empty, the prefetched content should be made
                  available to the build.
                name: PREFETCH_INPUT
                type: string
              - default: ''
                description: >-
                  Delete image tag after specified time. Empty means to keep the
                  image tag. Time values could be something like 1h, 2d, 3w for
                  hours, days, and weeks, respectively.
                name: IMAGE_EXPIRES_AFTER
                type: string
              - default: ''
                description: The image is built from this commit.
                name: COMMIT_SHA
                type: string
            results:
              - description: Digest of the image just built
                name: IMAGE_DIGEST
                type: string
              - description: Image repository where the built image was pushed
                name: IMAGE_URL
                type: string
              - description: Digests of the base images used for build
                name: BASE_IMAGES_DIGESTS
                type: string
              - description: The counting of Java components by publisher in JSON format
                name: SBOM_JAVA_COMPONENTS_COUNT
                type: string
              - description: >-
                  The Java dependencies that came from community sources such as
                  Maven central.
                name: JAVA_COMMUNITY_DEPENDENCIES
                type: string
            stepTemplate:
              env:
                - name: BUILDAH_FORMAT
                  value: oci
                - name: STORAGE_DRIVER
                  value: vfs
                - name: HERMETIC
                  value: 'false'
                - name: PREFETCH_INPUT
                - name: CONTEXT
                  value: .
                - name: DOCKERFILE
                  value: >-
                    https://raw.githubusercontent.com/devfile-samples/devfile-sample-java-springboot-basic/main/docker/Dockerfile
                - name: IMAGE
                  value: >-
                    quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                - name: TLSVERIFY
                  value: 'true'
                - name: IMAGE_EXPIRES_AFTER
                  value: 5d
              name: ''
              resources: {}
            steps:
              - env:
                  - name: COMMIT_SHA
                    value: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                image: 'quay.io/redhat-appstudio/buildah:v1.28'
                name: build
                resources:
                  limits:
                    cpu: '2'
                    memory: 4Gi
                  requests:
                    cpu: 250m
                    memory: 512Mi
                script: >
                  if [ -e "$CONTEXT/$DOCKERFILE" ]; then
                    dockerfile_path="$CONTEXT/$DOCKERFILE"
                  elif [ -e "$DOCKERFILE" ]; then
                    dockerfile_path="$DOCKERFILE"
                  elif echo "$DOCKERFILE" | grep -q "^https\?://"; then
                    echo "Fetch Dockerfile from $DOCKERFILE"
                    dockerfile_path=$(mktemp --suffix=-Dockerfile)
                    http_code=$(curl -s -L -w "%{http_code}" --output "$dockerfile_path" "$DOCKERFILE")
                    if [ $http_code != 200 ]; then
                      echo "No Dockerfile is fetched. Server responds $http_code"
                      exit 1
                    fi
                    http_code=$(curl -s -L -w "%{http_code}" --output "$dockerfile_path.dockerignore.tmp" "$DOCKERFILE.dockerignore")
                    if [ $http_code = 200 ]; then
                      echo "Fetched .dockerignore from $DOCKERFILE.dockerignore"
                      mv "$dockerfile_path.dockerignore.tmp" $CONTEXT/.dockerignore
                    fi
                  else
                    echo "Cannot find Dockerfile $DOCKERFILE"
                    exit 1
                  fi

                  if [ -n "$JVM_BUILD_WORKSPACE_ARTIFACT_CACHE_PORT_80_TCP_ADDR"
                  ] && grep -q '^\s*RUN \(./\)\?mvn' "$dockerfile_path"; then
                    sed -i -e "s|^\s*RUN \(\(./\)\?mvn\(.*\)\)|RUN echo \"<settings><mirrors><mirror><id>mirror.default</id><url>http://$JVM_BUILD_WORKSPACE_ARTIFACT_CACHE_PORT_80_TCP_ADDR/v1/cache/default/0/</url><mirrorOf>*</mirrorOf></mirror></mirrors></settings>\" > /tmp/settings.yaml; \1 -s /tmp/settings.yaml|g" "$dockerfile_path"
                    touch /var/lib/containers/java
                  fi


                  # Fixing group permission on /var/lib/containers

                  chown root:root /var/lib/containers


                  sed -i 's/^\s*short-name-mode\s*=\s*.*/short-name-mode =
                  "disabled"/' /etc/containers/registries.conf


                  # Setting new namespace to run buildah - 2^32-2

                  echo 'root:1:4294967294' | tee -a /etc/subuid >> /etc/subgid


                  if [ "${HERMETIC}" == "true" ]; then
                    BUILDAH_ARGS="--pull=never"
                    UNSHARE_ARGS="--net"
                    for image in $(grep -i '^\s*FROM' "$dockerfile_path" | sed 's/--platform=\S*//' | awk '{print $2}'); do
                      unshare -Ufp --keep-caps -r --map-users 1,1,65536 --map-groups 1,1,65536 -- buildah pull $image
                    done
                    echo "Build will be executed with network isolation"
                  fi


                  if [ -n "${PREFETCH_INPUT}" ]; then
                    mv cachi2 /tmp/
                    chmod -R go+rwX /tmp/cachi2
                    VOLUME_MOUNTS="--volume /tmp/cachi2:/cachi2"
                    sed -i 's|^\s*run |RUN . /cachi2/cachi2.env \&\& \\\n    |i' "$dockerfile_path"
                    echo "Prefetched content will be made available"
                  fi


                  LABELS=(
                    "--label" "build-date=$(date -u +'%Y-%m-%dT%H:%M:%S')"
                    "--label" "architecture=$(uname -m)"
                    "--label" "vcs-type=git"
                  )

                  [ -n "$COMMIT_SHA" ] && LABELS+=("--label"
                  "vcs-ref=$COMMIT_SHA")

                  [ -n "$IMAGE_EXPIRES_AFTER" ] && LABELS+=("--label"
                  "quay.expires-after=$IMAGE_EXPIRES_AFTER")


                  unshare -Uf $UNSHARE_ARGS --keep-caps -r --map-users 1,1,65536
                  --map-groups 1,1,65536 -- buildah build \
                    $VOLUME_MOUNTS \
                    $BUILDAH_ARGS \
                    ${LABELS[@]} \
                    --tls-verify=$TLSVERIFY --no-cache \
                    --ulimit nofile=4096:4096 \
                    -f "$dockerfile_path" -t $IMAGE $CONTEXT

                  container=$(buildah from --pull-never $IMAGE)

                  buildah mount $container | tee /workspace/container_path

                  echo $container > /workspace/container_name


                  # Save the SBOM produced by Cachi2 so it can be merged into
                  the final SBOM later

                  if [ -n "${PREFETCH_INPUT}" ]; then
                    cp /tmp/cachi2/output/bom.json ./sbom-cachi2.json
                  fi
                securityContext:
                  capabilities:
                    add:
                      - SETFCAP
                volumeMounts:
                  - mountPath: /var/lib/containers
                    name: varlibcontainers
                workingDir: /workspace/source
              - image: 'quay.io/redhat-appstudio/syft:v0.85.0'
                name: sbom-syft-generate
                resources: {}
                script: >
                  syft dir:/workspace/source
                  --file=/workspace/source/sbom-source.json
                  --output=cyclonedx-json

                  find $(cat /workspace/container_path) -xtype l -delete

                  syft dir:$(cat /workspace/container_path)
                  --file=/workspace/source/sbom-image.json
                  --output=cyclonedx-json
                volumeMounts:
                  - mountPath: /var/lib/containers
                    name: varlibcontainers
              - image: >-
                  quay.io/redhat-appstudio/hacbs-jvm-build-request-processor:1d417e6f1f3e68c6c537333b5759796eddae0afc
                name: analyse-dependencies-java-sbom
                resources: {}
                script: |
                  if [ -f /var/lib/containers/java ]; then
                    /opt/jboss/container/java/run/run-java.sh analyse-dependencies path $(cat /workspace/container_path) -s /workspace/source/sbom-image.json --task-run-name human-resources-on-pull-request-rlrj8-build-container --publishers /tekton/results/SBOM_JAVA_COMPONENTS_COUNT
                    sed -i 's/^/ /' /tekton/results/SBOM_JAVA_COMPONENTS_COUNT # Workaround for SRVKP-2875
                  else
                    touch /tekton/results/JAVA_COMMUNITY_DEPENDENCIES
                  fi
                securityContext:
                  runAsUser: 0
                volumeMounts:
                  - mountPath: /var/lib/containers
                    name: varlibcontainers
              - image: 'registry.access.redhat.com/ubi9/python-39:1-133.1692772345'
                name: merge-syft-sboms
                resources: {}
                script: >
                  #!/bin/python3

                  import json


                  # load SBOMs

                  with open("./sbom-image.json") as f:
                    image_sbom = json.load(f)

                  with open("./sbom-source.json") as f:
                    source_sbom = json.load(f)

                  # fetch unique components from available SBOMs

                  def get_identifier(component):
                    return component["name"] + '@' + component.get("version", "")

                  existing_components = [get_identifier(component) for component
                  in image_sbom["components"]]


                  for component in source_sbom["components"]:
                    if get_identifier(component) not in existing_components:
                      image_sbom["components"].append(component)
                      existing_components.append(get_identifier(component))

                  image_sbom["components"].sort(key=lambda c: get_identifier(c))


                  # write the CycloneDX unified SBOM

                  with open("./sbom-cyclonedx.json", "w") as f:
                    json.dump(image_sbom, f, indent=4)
                securityContext:
                  runAsUser: 0
                workingDir: /workspace/source
              - image: >-
                  quay.io/redhat-appstudio/cachi2:0.3.0@sha256:46097f22b57e4d48a3fce96d931e08ccfe3a3e6421362d5f9353961279078eef
                name: merge-cachi2-sbom
                resources: {}
                script: |
                  if [ -n "${PREFETCH_INPUT}" ]; then
                    echo "Merging contents of sbom-cachi2.json into sbom-cyclonedx.json"
                    /src/utils/merge_syft_sbom.py sbom-cachi2.json sbom-cyclonedx.json > sbom-temp.json
                    mv sbom-temp.json sbom-cyclonedx.json
                  else
                    echo "Skipping step since no Cachi2 SBOM was produced"
                  fi
                securityContext:
                  runAsUser: 0
                workingDir: /workspace/source
              - image: 'registry.access.redhat.com/ubi9/python-39:1-133.1692772345'
                name: create-purl-sbom
                resources: {}
                script: >
                  #!/bin/python3

                  import json


                  with open("./sbom-cyclonedx.json") as f:
                    cyclonedx_sbom = json.load(f)

                  purls = [{"purl": component["purl"]} for component in
                  cyclonedx_sbom["components"] if "purl" in component]

                  purl_content = {"image_contents": {"dependencies": purls}}


                  with open("sbom-purl.json", "w") as output_file:
                    json.dump(purl_content, output_file, indent=4)
                securityContext:
                  runAsUser: 0
                workingDir: /workspace/source
              - image: >-
                  registry.access.redhat.com/ubi9/buildah:9.0.0-19@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
                name: inject-sbom-and-push
                resources: {}
                script: >
                  # Expose base image digests

                  buildah images --format '{{ .Name }}:{{ .Tag }}@{{ .Digest }}'
                  | grep -v $IMAGE > /tekton/results/BASE_IMAGES_DIGESTS


                  base_image_name=$(buildah inspect --format '{{ index
                  .ImageAnnotations "org.opencontainers.image.base.name"}}'
                  $IMAGE | cut -f1 -d'@')

                  base_image_digest=$(buildah inspect --format '{{ index
                  .ImageAnnotations "org.opencontainers.image.base.digest"}}'
                  $IMAGE)

                  container=$(buildah from --pull-never $IMAGE)

                  buildah copy $container sbom-cyclonedx.json sbom-purl.json
                  /root/buildinfo/content_manifests/

                  buildah config -a
                  org.opencontainers.image.base.name=${base_image_name} -a
                  org.opencontainers.image.base.digest=${base_image_digest}
                  $container

                  buildah commit $container $IMAGE


                  status=-1

                  max_run=5

                  sleep_sec=10

                  for run in $(seq 1 $max_run); do
                    status=0
                    [ "$run" -gt 1 ] && sleep $sleep_sec
                    echo "Pushing sbom image to registry"
                    buildah push \
                      --tls-verify=$TLSVERIFY \
                      --digestfile /workspace/source/image-digest $IMAGE \
                      docker://$IMAGE && break || status=$?
                  done

                  if [ "$status" -ne 0 ]; then
                      echo "Failed to push sbom image to registry after ${max_run} tries"
                      exit 1
                  fi


                  cat "/workspace/source"/image-digest | tee
                  /tekton/results/IMAGE_DIGEST

                  echo -n "$IMAGE" | tee /tekton/results/IMAGE_URL
                securityContext:
                  capabilities:
                    add:
                      - SETFCAP
                  runAsUser: 0
                volumeMounts:
                  - mountPath: /var/lib/containers
                    name: varlibcontainers
                workingDir: /workspace/source
              - args:
                  - attach
                  - sbom
                  - '--sbom'
                  - sbom-cyclonedx.json
                  - '--type'
                  - cyclonedx
                  - >-
                    quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                image: 'quay.io/redhat-appstudio/cosign:v2.1.1'
                name: upload-sbom
                resources: {}
                workingDir: /workspace/source
            volumes:
              - emptyDir: {}
                name: varlibcontainers
            workspaces:
              - description: Workspace containing the source code to build.
                name: source
          duration: 1m 30s
          reason: Succeeded
      - name: inspect-image
        params:
          - name: IMAGE_URL
            value: $(tasks.build-container.results.IMAGE_URL)
          - name: IMAGE_DIGEST
            value: $(tasks.build-container.results.IMAGE_DIGEST)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: inspect-image
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-inspect-image:0.1@sha256:bbc286f0a2ad94e671ceb9d0f1debd96f36b8c38c1147c5030957820b4125fc6
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'false'
            operator: in
            values:
              - 'false'
        workspaces:
          - name: source
            workspace: workspace
        status:
          completionTime: '2023-08-25T14:58:48Z'
          conditions:
            - lastTransitionTime: '2023-08-25T14:58:48Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: human-resources-on-pull-request-rlrj8-inspect-image-pod
          startTime: '2023-08-25T14:58:25Z'
          steps:
            - container: step-inspect-image
              imageID: >-
                quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
              name: inspect-image
              terminated:
                containerID: >-
                  cri-o://8d1abe9b4cb1cda026b9da353d86abbdce8fc96cbc1029c7d1b601ed218f8f01
                exitCode: 0
                finishedAt: '2023-08-25T14:58:48Z'
                message: >-
                  [{"key":"BASE_IMAGE","value":"registry.access.redhat.com/ubi8/openjdk-17-runtime@sha256:14de89e89efc97aee3b50141108b7833708c3a93ad90bf89940025ab5267ba86","type":1},{"key":"BASE_IMAGE_REPOSITORY","value":"ubi8/openjdk-17-runtime","type":1},{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975528\",\"note\":\"Task
                  inspect-image completed: Check inspected JSON files under
                  /workspace/source/hacbs/inspect-image.\",\"namespace\":\"default\",\"successes\":1,\"failures\":0,\"warnings\":0}\n","type":1}]
                reason: Completed
                startedAt: '2023-08-25T14:58:47Z'
          taskResults:
            - name: BASE_IMAGE
              type: string
              value: >-
                registry.access.redhat.com/ubi8/openjdk-17-runtime@sha256:14de89e89efc97aee3b50141108b7833708c3a93ad90bf89940025ab5267ba86
            - name: BASE_IMAGE_REPOSITORY
              type: string
              value: ubi8/openjdk-17-runtime
            - name: TEST_OUTPUT
              type: string
              value: >
                {"result":"SUCCESS","timestamp":"1692975528","note":"Task
                inspect-image completed: Check inspected JSON files under
                /workspace/source/hacbs/inspect-image.","namespace":"default","successes":1,"failures":0,"warnings":0}
          taskSpec:
            description: >-
              Inspects and analyzes manifest data of the container's source
              image, and its base image (if available) using Skopeo. An image's
              manifest data contains information about the layers that make up
              the image, the platforms for which the image is intended, and
              other metadata about the image.
            params:
              - description: Fully qualified image name.
                name: IMAGE_URL
                type: string
              - description: Image digest.
                name: IMAGE_DIGEST
                type: string
              - default: ''
                description: 'unused, should be removed in next task version'
                name: DOCKER_AUTH
                type: string
            results:
              - description: Base image source image is built from.
                name: BASE_IMAGE
                type: string
              - description: Base image repository URL.
                name: BASE_IMAGE_REPOSITORY
                type: string
              - description: Tekton task test output.
                name: TEST_OUTPUT
                type: string
            steps:
              - env:
                  - name: IMAGE_URL
                    value: >-
                      quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                  - name: IMAGE_DIGEST
                    value: >-
                      sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d
                image: >-
                  quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
                name: inspect-image
                resources: {}
                script: >
                  #!/usr/bin/env bash

                  source /utils.sh

                  IMAGE_INSPECT=image_inspect.json

                  BASE_IMAGE_INSPECT=base_image_inspect.json

                  RAW_IMAGE_INSPECT=raw_image_inspect.json


                  IMAGE_URL="${IMAGE_URL}@${IMAGE_DIGEST}"

                  # Given a tag and a the digest in the IMAGE_URL we opt to use
                  the digest alone

                  # this is because containers/image currently doesn't support
                  image references

                  # that contain both. See
                  https://github.com/containers/image/issues/1736

                  if [[ "${IMAGE_URL}" == *":"*"@"* ]]; then
                    IMAGE_URL="${IMAGE_URL/:*@/@}"
                  fi


                  status=-1

                  max_run=5

                  sleep_sec=10

                  for run in $(seq 1 $max_run); do
                    status=0
                    [ "$run" -gt 1 ] && sleep $sleep_sec  # skip last sleep
                    echo "Inspecting manifest for source image ${IMAGE_URL} (try $run/$max_run)."
                    skopeo inspect --no-tags docker://"${IMAGE_URL}" > $IMAGE_INSPECT && break || status=$?
                  done

                  if [ "$status" -ne 0 ]; then
                      echo "Failed to inspect image ${IMAGE_URL}"
                      note="Task inspect-image failed: Encountered errors while inspecting image. For details, check Tekton task log."
                      TEST_OUTPUT=$(make_result_json -r ERROR -t "$note")
                      echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
                      exit 0
                  fi

                  echo "Image ${IMAGE_URL} metadata:"

                  cat "$IMAGE_INSPECT"


                  for run in $(seq 1 $max_run); do
                    status=0
                    [ "$run" -gt 1 ] && sleep $sleep_sec  # skip last sleep
                    echo "Inspecting raw image manifest ${IMAGE_URL} (try $run/$max_run)."
                    skopeo inspect --no-tags --raw docker://"${IMAGE_URL}" > $RAW_IMAGE_INSPECT && break || status=$?
                  done

                  if [ "$status" -ne 0 ]; then
                      echo "Failed to get raw metadata of image ${IMAGE_URL}"
                      note="Task inspect-image failed: Encountered errors while inspecting image. For details, check Tekton task log."
                      TEST_OUTPUT=$(make_result_json -r ERROR -t "$note")
                      echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
                      exit 0
                  fi

                  echo "Image ${IMAGE_URL} raw metadata:"

                  cat "$RAW_IMAGE_INSPECT" | jq  # jq for readable formatting


                  echo "Getting base image manifest for source image
                  ${IMAGE_URL}."

                  BASE_IMAGE_NAME="$(jq -r
                  ".annotations.\"org.opencontainers.image.base.name\""
                  $RAW_IMAGE_INSPECT)"

                  BASE_IMAGE_DIGEST="$(jq -r
                  ".annotations.\"org.opencontainers.image.base.digest\""
                  $RAW_IMAGE_INSPECT)"

                  if [ $BASE_IMAGE_NAME == 'null' ]; then
                    echo "Cannot get base image info from annotations."
                    BASE_IMAGE_NAME="$(jq -r ".Labels.\"org.opencontainers.image.base.name\"" $IMAGE_INSPECT)"
                    BASE_IMAGE_DIGEST="$(jq -r ".annotations.\"org.opencontainers.image.base.digest\"" $IMAGE_INSPECT)"
                    if [ "$BASE_IMAGE_NAME" == 'null' ]; then
                      echo "Cannot get base image info from Labels. For details, check source image ${IMAGE_URL}."
                      exit 0
                    fi
                  fi

                  if [ -z "$BASE_IMAGE_NAME" ]; then
                    echo "Source image ${IMAGE_URL} is built from scratch, so there is no base image."
                    exit 0
                  fi


                  BASE_IMAGE="${BASE_IMAGE_NAME%:*}@$BASE_IMAGE_DIGEST"

                  echo "Detected base image: $BASE_IMAGE"

                  echo -n "$BASE_IMAGE" > /tekton/results/BASE_IMAGE


                  for run in $(seq 1 $max_run); do
                    status=0
                    [ "$run" -gt 1 ] && sleep $sleep_sec  # skip last sleep
                    echo "Inspecting base image ${BASE_IMAGE} (try $run/$max_run)."
                    skopeo inspect --no-tags "docker://$BASE_IMAGE"  > $BASE_IMAGE_INSPECT && break || status=$?
                  done

                  if [ "$status" -ne 0 ]; then
                      echo "Failed to inspect base image ${BASE_IMAGE}"
                      note="Task inspect-image failed: Encountered errors while inspecting image. For details, check Tekton task log."
                      TEST_OUTPUT=$(make_result_json -r ERROR -t "$note")
                      echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
                      exit 0
                  fi


                  BASE_IMAGE_REPOSITORY="$(jq -r '.Name | sub("[^/]+/"; "") |
                  sub("[:@].*"; "")' "$BASE_IMAGE_INSPECT")"

                  echo "Detected base image repository: $BASE_IMAGE_REPOSITORY"

                  echo -n "$BASE_IMAGE_REPOSITORY" >
                  /tekton/results/BASE_IMAGE_REPOSITORY


                  note="Task inspect-image completed: Check inspected JSON files
                  under /workspace/source/hacbs/inspect-image."

                  TEST_OUTPUT=$(make_result_json -r SUCCESS -s 1 -t "$note")

                  echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
                securityContext:
                  capabilities:
                    add:
                      - SETFCAP
                  runAsUser: 0
                workingDir: /workspace/source/hacbs/inspect-image
            workspaces:
              - name: source
          duration: 23s
          reason: Succeeded
      - name: label-check
        runAfter:
          - inspect-image
        taskRef:
          kind: Task
          params:
            - name: name
              value: label-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-label-check:0.1@sha256:0c0739fdda24cd1e3587bbab9b07d4493efc21884baac7723f4b446e95bf1fd3
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'false'
            operator: in
            values:
              - 'false'
        workspaces:
          - name: workspace
            workspace: workspace
        status:
          completionTime: '2023-08-25T14:59:15Z'
          conditions:
            - lastTransitionTime: '2023-08-25T14:59:15Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: human-resources-on-pull-request-rlrj8-label-check-pod
          startTime: '2023-08-25T14:58:49Z'
          steps:
            - container: step-surface-level-checks-required-labels
              imageID: >-
                quay.io/redhat-appstudio/hacbs-test@sha256:82b43bffe4eacc717239424f64478b18f36528df47c2d11df3a8d031e81a3c67
              name: surface-level-checks-required-labels
              terminated:
                containerID: >-
                  cri-o://446c06967a1115140d01c2c34bb74cdf00ba5ca6af970d8891c796f9d4b762b6
                exitCode: 0
                finishedAt: '2023-08-25T14:59:14Z'
                message: >-
                  [{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975554\",\"note\":\"For
                  details, check Tekton task
                  log.\",\"namespace\":\"required_checks\",\"successes\":21,\"failures\":0,\"warnings\":0}\n","type":1}]
                reason: Completed
                startedAt: '2023-08-25T14:59:14Z'
          taskResults:
            - name: TEST_OUTPUT
              type: string
              value: >
                {"result":"SUCCESS","timestamp":"1692975554","note":"For
                details, check Tekton task
                log.","namespace":"required_checks","successes":21,"failures":0,"warnings":0}
          taskSpec:
            description: >-
              Verifies whether an image contains the best practice labels using
              Conftest.
            params:
              - default: /project/image/
                description: Path to directory containing Conftest policies.
                name: POLICY_DIR
                type: string
              - default: required_checks
                description: Namespace for Conftest policy.
                name: POLICY_NAMESPACE
                type: string
            results:
              - description: Tekton task test output.
                name: TEST_OUTPUT
                type: string
            steps:
              - env:
                  - name: POLICY_NAMESPACE
                    value: required_checks
                  - name: POLICY_DIR
                    value: /project/image/
                image: >-
                  quay.io/redhat-appstudio/hacbs-test:v1.1.0@sha256:82b43bffe4eacc717239424f64478b18f36528df47c2d11df3a8d031e81a3c67
                name: surface-level-checks-required-labels
                resources: {}
                script: >
                  #!/usr/bin/env bash


                  . /utils.sh

                  if [ ! -s ../inspect-image/image_inspect.json ]; then
                    echo "File $(workspaces.source.path)/hacbs/inspect-image/image_inspect.json did not generate correctly. Check task inspect-image log."
                    note="Task label-check failed: $(workspaces.source.path)/hacbs/inspect-image/image_inspect.json did not generate correctly. For details, check Tekton task result TEST_OUTPUT in task inspect-image."
                    TEST_OUTPUT=$(make_result_json -r ERROR -t "$note")
                    echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
                    exit 0
                  fi


                  CONFTEST_OPTIONS=""

                  if [ -s "../inspect-image/base_image_inspect.json" ]; then
                    CONFTEST_OPTIONS="-d=../inspect-image/base_image_inspect.json"
                  fi


                  echo "Running conftest using $POLICY_DIR policy,
                  $POLICY_NAMESPACE namespace."

                  /usr/bin/conftest test --no-fail
                  ../inspect-image/image_inspect.json "${CONFTEST_OPTIONS}" \

                  --policy $POLICY_DIR --namespace $POLICY_NAMESPACE \

                  --output=json 2> stderr.txt | tee label_check_output.json


                  if [ ! -z $(cat stderr.txt) ]; then
                    echo "label-check test encountered the following error:"
                    cat stderr.txt
                    note="Task label-check failed: Command conftest failed. For details, check Tekton task log."
                    ERROR_OUTPUT=$(make_result_json -r "ERROR" -t "$note")
                  fi


                  TEST_OUTPUT=

                  parse_test_output label-check conftest label_check_output.json
                  || true


                  echo "${TEST_OUTPUT:-${ERROR_OUTPUT}}" | tee
                  /tekton/results/TEST_OUTPUT
                securityContext:
                  capabilities:
                    add:
                      - SETFCAP
                workingDir: /workspace/workspace/hacbs/label-check-required_checks
            workspaces:
              - name: workspace
          duration: 26s
          reason: Succeeded
      - name: optional-label-check
        params:
          - name: POLICY_NAMESPACE
            value: optional_checks
        runAfter:
          - inspect-image
        taskRef:
          kind: Task
          params:
            - name: name
              value: label-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-label-check:0.1@sha256:0c0739fdda24cd1e3587bbab9b07d4493efc21884baac7723f4b446e95bf1fd3
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'true'
            operator: in
            values:
              - 'false'
        workspaces:
          - name: workspace
            workspace: workspace
        status:
          reason: Skipped
      - name: deprecated-base-image-check
        params:
          - name: BASE_IMAGES_DIGESTS
            value: $(tasks.build-container.results.BASE_IMAGES_DIGESTS)
        taskRef:
          kind: Task
          params:
            - name: name
              value: deprecated-image-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-deprecated-image-check:0.2@sha256:58d16de95b4ca597f7f860fb85d6206e549910fa7a8d2a2cc229558f791ad329
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'false'
            operator: in
            values:
              - 'false'
        workspaces:
          - name: test-ws
            workspace: workspace
        status:
          completionTime: '2023-08-25T14:58:48Z'
          conditions:
            - lastTransitionTime: '2023-08-25T14:58:48Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: humeaa4363f279f84d4e98a2feefb857604072fb4347373bf5b615f86ad-pod
          startTime: '2023-08-25T14:58:25Z'
          steps:
            - container: step-query-pyxis
              imageID: >-
                registry.access.redhat.com/ubi8/ubi-minimal@sha256:7394c071ed74ace08cfd51f881c94067fa7a570e7f7e4a0ef0aff1b4f6a2a949
              name: query-pyxis
              terminated:
                containerID: >-
                  cri-o://521ca0e70485705fdf7b7b7b15e517019a6099ebb6099f41128d980f054ea3e6
                exitCode: 0
                finishedAt: '2023-08-25T14:58:47Z'
                message: >-
                  [{"key":"PYXIS_HTTP_CODE","value":"200
                  registry.access.redhat.com ubi8/openjdk-17\n200
                  registry.access.redhat.com
                  ubi8/openjdk-17-runtime\n","type":1}]
                reason: Completed
                startedAt: '2023-08-25T14:58:47Z'
            - container: step-run-conftest
              imageID: >-
                quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
              name: run-conftest
              terminated:
                containerID: >-
                  cri-o://8f8de37355b5e0fa6d6fbc0f55ed28dc2cd9ad600dc6cc27a62a45434fb4ecb3
                exitCode: 0
                finishedAt: '2023-08-25T14:58:48Z'
                message: >-
                  [{"key":"PYXIS_HTTP_CODE","value":"200
                  registry.access.redhat.com ubi8/openjdk-17\n200
                  registry.access.redhat.com
                  ubi8/openjdk-17-runtime\n","type":1},{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975528\",\"note\":\"Task
                  deprecated-image-check completed: Check result for task
                  result.\",\"namespace\":\"required_checks\",\"successes\":2,\"failures\":0,\"warnings\":0}\n","type":1}]
                reason: Completed
                startedAt: '2023-08-25T14:58:48Z'
          taskResults:
            - name: PYXIS_HTTP_CODE
              type: string
              value: |
                200 registry.access.redhat.com ubi8/openjdk-17
                200 registry.access.redhat.com ubi8/openjdk-17-runtime
            - name: TEST_OUTPUT
              type: string
              value: >
                {"result":"SUCCESS","timestamp":"1692975528","note":"Task
                deprecated-image-check completed: Check result for task
                result.","namespace":"required_checks","successes":2,"failures":0,"warnings":0}
          taskSpec:
            description: >-
              Identifies the unmaintained and potentially insecure deprecated
              base images. Pyxis API collects metadata from image repository,
              and Conftest applies supplied policy to identify the deprecated
              images using that metadata.
            params:
              - default: /project/repository/
                description: Path to directory containing Conftest policies.
                name: POLICY_DIR
                type: string
              - default: required_checks
                description: Namespace for Conftest policy.
                name: POLICY_NAMESPACE
                type: string
              - description: Digests of base build images.
                name: BASE_IMAGES_DIGESTS
                type: string
            results:
              - description: HTTP code returned by Pyxis API endpoint.
                name: PYXIS_HTTP_CODE
                type: string
              - description: Tekton task test output.
                name: TEST_OUTPUT
                type: string
            steps:
              - env:
                  - name: BASE_IMAGES_DIGESTS
                    value: >
                      registry.access.redhat.com/ubi8/openjdk-17:1.15-1.1682053058@sha256:b00f687d913b8d1e027f7eabd6765de6c8d469629bef9550f10dbf207af24fe5

                      registry.access.redhat.com/ubi8/openjdk-17-runtime:1.15-1.1682053056@sha256:f921cf1f9147e4b306908f3bcb61dd215b4a51970f8db560ede02ee6a492fa99
                image: >-
                  registry.access.redhat.com/ubi8/ubi-minimal:8.8-1037@sha256:8d43664c250c72d35af8498c7ff76a9f0d42f16b9b3b29f0caa747121778de0e
                name: query-pyxis
                resources: {}
                script: >
                  #!/usr/bin/env bash

                  readarray -t IMAGE_ARRAY < <(echo -n "$BASE_IMAGES_DIGESTS" |
                  sed 's/\\n/\'$'\n''/g')

                  for BASE_IMAGE in ${IMAGE_ARRAY[@]};

                  do
                    IFS=:'/' read -r IMAGE_REGISTRY IMAGE_WITH_TAG <<< $BASE_IMAGE; echo "[$IMAGE_REGISTRY] [$IMAGE_WITH_TAG]"
                    IMAGE_REPOSITORY=echo $IMAGE_WITH_TAG | cut -d ":" -f1
                    IMAGE_REGISTRY=${IMAGE_REGISTRY//registry.redhat.io/registry.access.redhat.com}
                    export IMAGE_REPO_PATH=/workspace/test-ws/${IMAGE_REPOSITORY}
                    mkdir -p ${IMAGE_REPO_PATH}
                    echo "Querying Pyxis for $BASE_IMAGE."
                    http_code=$(curl -s -k -o ${IMAGE_REPO_PATH}/repository_data.json -w '%{http_code}' "https://catalog.redhat.com/api/containers/v1/repositories/registry/${IMAGE_REGISTRY}/repository/${IMAGE_REPOSITORY}")
                    echo "Response code: $http_code."
                    echo $http_code $IMAGE_REGISTRY $IMAGE_REPOSITORY>> /tekton/results/PYXIS_HTTP_CODE
                  done
              - env:
                  - name: POLICY_DIR
                    value: /project/repository/
                  - name: POLICY_NAMESPACE
                    value: required_checks
                image: >-
                  quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
                name: run-conftest
                resources: {}
                script: >
                  #!/usr/bin/env sh

                  source /utils.sh


                  success_counter=0

                  failure_counter=0

                  error_counter=0

                  if [ ! -f /tekton/results/PYXIS_HTTP_CODE ]; then
                    error_counter=$((error_counter++))
                  fi

                  while IFS= read -r line

                  do
                    IFS=:' ' read -r http_code IMAGE_REGISTRY IMAGE_REPOSITORY <<< $line; echo "[$http_code] [$IMAGE_REGISTRY] [$IMAGE_REPOSITORY]"
                    export IMAGE_REPO_PATH=/workspace/test-ws/${IMAGE_REPOSITORY}
                    if [ "$http_code" == "200" ];
                    then
                      echo "Running conftest using $POLICY_DIR policy, $POLICY_NAMESPACE namespace."
                      /usr/bin/conftest test --no-fail ${IMAGE_REPO_PATH}/repository_data.json \
                      --policy $POLICY_DIR --namespace $POLICY_NAMESPACE \
                      --output=json 2> ${IMAGE_REPO_PATH}/stderr.txt | tee ${IMAGE_REPO_PATH}/deprecated_image_check_output.json

                      failure_counter=$((failure_counter+$(jq -r '.[].failures|length' ${IMAGE_REPO_PATH}/deprecated_image_check_output.json)))
                      success_counter=$((success_counter+$(jq -r '.[].successes' ${IMAGE_REPO_PATH}/deprecated_image_check_output.json)))

                    elif [ "$http_code" == "404" ];
                    then
                      echo "Registry/image ${IMAGE_REGISTRY}/${IMAGE_REPOSITORY} not found in Pyxis." >> /workspace/test-ws/stderr.txt
                      cat /workspace/test-ws/stderr.txt
                    else
                      echo "Unexpected error HTTP code $http_code) occurred for registry/image ${IMAGE_REGISTRY}/${IMAGE_REPOSITORY}." >> /workspace/test-ws/stderr.txt
                      cat /workspace/test-ws/stderr.txt
                      error_counter=$((error_counter++))
                      exit 0
                    fi
                  done < /tekton/results/PYXIS_HTTP_CODE


                  note="Task deprecated-image-check failed: Command conftest
                  failed. For details, check Tekton task log."

                  ERROR_OUTPUT=$(make_result_json -r ERROR -n
                  "$POLICY_NAMESPACE" -t "$note")

                  if [[ "$error_counter" == 0 && "$success_counter" > 0 ]];

                  then
                    if [[ "${failure_counter}" -gt 0 ]]; then RES="FAILURE"; else RES="SUCCESS"; fi
                    note="Task deprecated-image-check completed: Check result for task result."
                    TEST_OUTPUT=$(make_result_json \
                      -r "${RES}" -n "$POLICY_NAMESPACE" \
                      -s "${success_counter}" -f "${failure_counter}" -t "$note")
                  fi

                  echo "${TEST_OUTPUT:-${ERROR_OUTPUT}}" | tee
                  /tekton/results/TEST_OUTPUT
            workspaces:
              - name: test-ws
          duration: 23s
          reason: Succeeded
      - name: clair-scan
        params:
          - name: image-digest
            value: $(tasks.build-container.results.IMAGE_DIGEST)
          - name: image-url
            value: $(tasks.build-container.results.IMAGE_URL)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: clair-scan
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-clair-scan:0.1@sha256:c5602d9d6dd797da98e98fde8471ea55a788c30f74f2192807910ce5436e9b66
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'false'
            operator: in
            values:
              - 'false'
        status:
          completionTime: '2023-08-25T14:59:41Z'
          conditions:
            - lastTransitionTime: '2023-08-25T14:59:41Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: human-resources-on-pull-request-rlrj8-clair-scan-pod
          startTime: '2023-08-25T14:59:25Z'
          steps:
            - container: step-get-vulnerabilities
              imageID: >-
                quay.io/redhat-appstudio/clair-in-ci@sha256:ff09557845e2ccb555fcce534e27053976260ebd11c984e3c06d2062bec336e1
              name: get-vulnerabilities
              terminated:
                containerID: >-
                  cri-o://3e1116ef2e6822f8ccafdd228050ad8331098267b1465f75be97fb2d9bd94a50
                exitCode: 0
                finishedAt: '2023-08-25T14:59:40Z'
                reason: Completed
                startedAt: '2023-08-25T14:59:30Z'
            - container: step-conftest-vulnerabilities
              imageID: >-
                quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
              name: conftest-vulnerabilities
              terminated:
                containerID: >-
                  cri-o://b74915dd54a8815d854d41bfd7853947e77e95cedf7de92ab5c746c142ad20e7
                exitCode: 0
                finishedAt: '2023-08-25T14:59:41Z'
                reason: Completed
                startedAt: '2023-08-25T14:59:41Z'
            - container: step-test-format-result
              imageID: >-
                quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
              name: test-format-result
              terminated:
                containerID: >-
                  cri-o://e11d42d22fef1982c79a6914661f4a02d2d8bab0d24ac54d82058a01dfe2388a
                exitCode: 0
                finishedAt: '2023-08-25T14:59:41Z'
                message: >-
                  [{"key":"CLAIR_SCAN_RESULT","value":"{\"vulnerabilities\":{\"critical\":0,\"high\":2,\"medium\":12,\"low\":3}}\n","type":1},{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975581\",\"note\":\"Task
                  clair-scan completed: Refer to Tekton task result
                  CLAIR_SCAN_RESULT for vulnerabilities scanned by
                  Clair.\",\"namespace\":\"default\",\"successes\":0,\"failures\":0,\"warnings\":0}\n","type":1}]
                reason: Completed
                startedAt: '2023-08-25T14:59:41Z'
          taskResults:
            - name: CLAIR_SCAN_RESULT
              type: string
              value: |
                {"vulnerabilities":{"critical":0,"high":2,"medium":12,"low":3}}
            - name: TEST_OUTPUT
              type: string
              value: >
                {"result":"SUCCESS","timestamp":"1692975581","note":"Task
                clair-scan completed: Refer to Tekton task result
                CLAIR_SCAN_RESULT for vulnerabilities scanned by
                Clair.","namespace":"default","successes":0,"failures":0,"warnings":0}
          taskSpec:
            description: >-
              Scans container images for vulnerabilities using Clair, by
              comparing the components of container image against Clair's
              vulnerability databases.
            params:
              - description: Image digest to scan.
                name: image-digest
                type: string
              - description: Image URL.
                name: image-url
                type: string
              - default: ''
                description: 'unused, should be removed in next task version.'
                name: docker-auth
                type: string
            results:
              - description: Tekton task test output.
                name: TEST_OUTPUT
                type: string
              - description: Clair scan result.
                name: CLAIR_SCAN_RESULT
                type: string
            steps:
              - env:
                  - name: IMAGE_URL
                    value: >-
                      quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                  - name: IMAGE_DIGEST
                    value: >-
                      sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d
                image: 'quay.io/redhat-appstudio/clair-in-ci:latest'
                imagePullPolicy: Always
                name: get-vulnerabilities
                resources: {}
                script: >
                  #!/usr/bin/env bash


                  imagewithouttag=$(echo $IMAGE_URL | sed "s/\(.*\):.*/\1/" | tr
                  -d '\n')

                  # strip new-line escape symbol from parameter and save it to
                  variable

                  imageanddigest=$(echo $imagewithouttag@$IMAGE_DIGEST)


                  clair-action report --image-ref=$imageanddigest
                  --db-path=/tmp/matcher.db --format=quay | tee
                  /tekton/home/clair-result.json || true
              - image: >-
                  quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
                name: conftest-vulnerabilities
                resources: {}
                script: |
                  if [ ! -s /tekton/home/clair-result.json ]; then
                    echo "Previous step [get-vulnerabilities] failed: /tekton/home/clair-result.json is empty."
                  else
                    /usr/bin/conftest test --no-fail /tekton/home/clair-result.json \
                    --policy /project/clair/vulnerabilities-check.rego --namespace required_checks \
                    --output=json | tee /tekton/home/clair-vulnerabilities.json || true
                  fi
                securityContext:
                  capabilities:
                    add:
                      - SETFCAP
              - image: >-
                  quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
                name: test-format-result
                resources: {}
                script: >
                  #!/usr/bin/env bash

                  . /utils.sh


                  if [[ ! -f /tekton/home/clair-vulnerabilities.json ]]; then
                    note="Task clair-scan failed: /tekton/home/clair-vulnerabilities.json did not generate. For details, check Tekton task log."
                    TEST_OUTPUT=$(make_result_json -r "ERROR" -t "$note")
                    echo "/tekton/home/clair-vulnerabilities.json did not generate correctly. For details, check conftest command in Tekton task log."
                    echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
                    exit 0
                  fi


                  jq -rce \
                    '{vulnerabilities:{
                        critical: (.[] | .warnings? // [] | map(select(.metadata.details.name=="clair_critical_vulnerabilities").metadata."vulnerabilities_number" // 0)| add // 0),
                        high: (.[] | .warnings? // [] | map(select(.metadata.details.name=="clair_high_vulnerabilities").metadata."vulnerabilities_number" // 0)| add // 0),
                        medium: (.[] | .warnings? // [] | map(select(.metadata.details.name=="clair_medium_vulnerabilities").metadata."vulnerabilities_number" // 0)| add // 0),
                        low: (.[] | .warnings? // [] | map(select(.metadata.details.name=="clair_low_vulnerabilities").metadata."vulnerabilities_number" // 0)| add // 0)
                      }}' /tekton/home/clair-vulnerabilities.json | tee /tekton/results/CLAIR_SCAN_RESULT

                  note="Task clair-scan completed: Refer to Tekton task result
                  CLAIR_SCAN_RESULT for vulnerabilities scanned by Clair."

                  TEST_OUTPUT=$(make_result_json -r "SUCCESS" -t "$note")

                  echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
          duration: 16s
          reason: Succeeded
      - name: sast-snyk-check
        params:
          - name: SNYK_SECRET
            value: ''
        runAfter:
          - clone-repository
        taskRef:
          kind: Task
          params:
            - name: name
              value: sast-snyk-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-sast-snyk-check:0.1@sha256:9dcd450b454705b9fe22c5f8f7bb7305cebc3cb73e783b85e047f7e721994189
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'false'
            operator: in
            values:
              - 'false'
          - input: ''
            operator: notin
            values:
              - ''
        workspaces:
          - name: workspace
            workspace: workspace
        status:
          reason: Skipped
      - name: clamav-scan
        params:
          - name: image-digest
            value: $(tasks.build-container.results.IMAGE_DIGEST)
          - name: image-url
            value: $(tasks.build-container.results.IMAGE_URL)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: clamav-scan
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-clamav-scan:0.1@sha256:cd4e301dd849cbdf7b8e38fd8f4915970b5b60174770df632a6b38ea93028d44
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'false'
            operator: in
            values:
              - 'false'
        status:
          completionTime: '2023-08-25T15:01:02Z'
          conditions:
            - lastTransitionTime: '2023-08-25T15:01:02Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: human-resources-on-pull-request-rlrj8-clamav-scan-pod
          sidecars:
            - container: sidecar-database
              imageID: >-
                quay.io/redhat-appstudio/clamav-db@sha256:703c928d5d34a6004f50e28301d4aa642d3eb18edaa6a697dc73fbe72c46ffe5
              name: database
              terminated:
                containerID: >-
                  cri-o://9105b69f157a8f5a92856ab8dd27c5d392d32ab7ad00c7442825f3112b20f04f
                exitCode: 0
                finishedAt: '2023-08-25T14:59:29Z'
                reason: Completed
                startedAt: '2023-08-25T14:59:29Z'
          startTime: '2023-08-25T14:59:25Z'
          steps:
            - container: step-extract-and-scan-image
              imageID: >-
                quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
              name: extract-and-scan-image
              terminated:
                containerID: >-
                  cri-o://16fa2bfc99dc510afdc9983de8501b57013101c128ba467905d118134ef80a0e
                exitCode: 0
                finishedAt: '2023-08-25T15:01:01Z'
                reason: Completed
                startedAt: '2023-08-25T14:59:31Z'
            - container: step-modify-clam-output-to-json
              imageID: >-
                quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
              name: modify-clam-output-to-json
              terminated:
                containerID: >-
                  cri-o://0dafda6e48f6d18bb0aa3421d0dc64804eea9c35fb2f2eeaf4d4e09f01ac83ea
                exitCode: 0
                finishedAt: '2023-08-25T15:01:02Z'
                reason: Completed
                startedAt: '2023-08-25T15:01:01Z'
            - container: step-store-hacbs-test-output-result
              imageID: >-
                quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
              name: store-hacbs-test-output-result
              terminated:
                containerID: >-
                  cri-o://24ae0f28387e6b1287857396af6be637a3c03a308e5b3d1396451630e1c1514c
                exitCode: 0
                finishedAt: '2023-08-25T15:01:02Z'
                message: >-
                  [{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975662\",\"note\":\"Task
                  clamav-scan completed: Check result for antivirus scan
                  result.\",\"namespace\":\"default\",\"successes\":1,\"failures\":0,\"warnings\":0}\n","type":1}]
                reason: Completed
                startedAt: '2023-08-25T15:01:02Z'
          taskResults:
            - name: TEST_OUTPUT
              type: string
              value: >
                {"result":"SUCCESS","timestamp":"1692975662","note":"Task
                clamav-scan completed: Check result for antivirus scan
                result.","namespace":"default","successes":1,"failures":0,"warnings":0}
          taskSpec:
            description: >-
              Scans the content of container images for viruses, malware, and
              other malicious content using ClamAV antivirus scanner.
            params:
              - description: Image digest to scan.
                name: image-digest
                type: string
              - description: Image URL.
                name: image-url
                type: string
              - default: ''
                description: unused
                name: docker-auth
                type: string
            results:
              - description: Tekton task test output.
                name: TEST_OUTPUT
                type: string
            sidecars:
              - image: 'quay.io/redhat-appstudio/clamav-db:v1'
                imagePullPolicy: Always
                name: database
                resources: {}
                script: |
                  #!/usr/bin/env bash
                  clamscan --version
                  cp -r /var/lib/clamav/* /tmp/clamdb
                volumeMounts:
                  - mountPath: /tmp/clamdb
                    name: dbfolder
            steps:
              - env:
                  - name: HOME
                    value: /work
                  - name: IMAGE_URL
                    value: >-
                      quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                  - name: IMAGE_DIGEST
                    value: >-
                      sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d
                image: >-
                  quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
                name: extract-and-scan-image
                resources:
                  limits:
                    cpu: '2'
                    memory: 4Gi
                  requests:
                    cpu: 10m
                    memory: 512Mi
                script: >
                  imagewithouttag=$(echo $IMAGE_URL | sed "s/\(.*\):.*/\1/" | tr
                  -d '\n')


                  # strip new-line escape symbol from parameter and save it to
                  variable

                  imageanddigest=$(echo $imagewithouttag@$IMAGE_DIGEST)


                  # check if image is attestation one, skip the clamav scan in
                  such case

                  if [[ $imageanddigest == *.att ]]

                  then
                      echo "$imageanddigest is an attestation image. Skipping ClamAV scan."
                      exit 0
                  fi

                  mkdir content

                  cd content

                  echo Extracting image.

                  if ! oc image extract --registry-config ~/.docker/config.json
                  $imageanddigest; then
                    echo "Unable to extract image. Skipping ClamAV scan!"
                    exit 0
                  fi

                  echo Extraction done.

                  clamscan -ri --max-scansize=250M | tee
                  /tekton/home/clamscan-result.log

                  echo "Executed-on: Scan was executed on version - $(clamscan
                  --version)" | tee -a /tekton/home/clamscan-result.log
                securityContext:
                  runAsUser: 1000
                volumeMounts:
                  - mountPath: /var/lib/clamav
                    name: dbfolder
                  - mountPath: /work
                    name: work
                workingDir: /work
              - image: >-
                  quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
                name: modify-clam-output-to-json
                resources: {}
                script: >
                  #!/usr/bin/env python3.9

                  import json

                  import dateutil.parser as parser

                  import os


                  clamscan_result = "/tekton/home/clamscan-result.log"

                  if not os.path.exists(clamscan_result) or
                  os.stat(clamscan_result).st_size == 0:
                      print("clamscan-result.log file is empty, so compiled code not extracted. Parsing skipped.")
                      exit(0)

                  with open(clamscan_result, "r") as file:
                      clam_result_str = file.read()

                  def clam_result_str_to_json(clam_result_str):

                      clam_result_list = clam_result_str.split("\n")
                      clam_result_list.remove('')

                      results_marker = \
                          clam_result_list.index("----------- SCAN SUMMARY -----------")

                      hit_list = clam_result_list[:results_marker]
                      summary_list = clam_result_list[(results_marker + 1):]

                      r_dict = { "hits": hit_list }
                      for item in summary_list:
                          # in case of blank lines
                          if not item:
                              continue
                          split_index = [c == ':' for c in item].index(True)
                          key = item[:split_index].lower()
                          key = key.replace(" ", "_")
                          value = item[(split_index + 1):].strip(" ")
                          if (key == "start_date" or key == "end_date"):
                            isodate = parser.parse(value)
                            value = isodate.isoformat()
                          r_dict[key] = value
                      print(json.dumps(r_dict))
                      with open('/tekton/home/clamscan-result.json', 'w') as f:
                        print(json.dumps(r_dict), file=f)

                  def main():
                      clam_result_str_to_json(clam_result_str)

                  if __name__ == "__main__":
                      main()
              - image: >-
                  quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
                name: store-hacbs-test-output-result
                resources: {}
                script: >
                  #!/usr/bin/env bash

                  source /utils.sh


                  if [ -f /tekton/home/clamscan-result.json ];

                  then
                    cat /tekton/home/clamscan-result.json
                    INFECTED_FILES=$(jq -r '.infected_files' /tekton/home/clamscan-result.json || true )
                    if [ -z "${INFECTED_FILES}" ]; then
                      echo "Failed to get number of infected files."
                      note="Task clamav-scan failed: Unable to get number of infected files from /tekton/home/clamscan-result.json. For details, check Tekton task log."
                    else
                      if [[ "${INFECTED_FILES}" -gt 0 ]]; then RES="FAILURE"; else RES="SUCCESS"; fi
                      note="Task clamav-scan completed: Check result for antivirus scan result."
                      TEST_OUTPUT=$(make_result_json -r "${RES}" -s 1 -f "${INFECTED_FILES}" -t "$note")
                    fi
                  else
                    note="Task clamav-scan failed: /tekton/home/clamscan-result.json doesn't exist. For details, check Tekton task log."
                  fi


                  ERROR_OUTPUT=$(make_result_json -r "ERROR" -t "$note")

                  echo "${TEST_OUTPUT:-${ERROR_OUTPUT}}" | tee
                  /tekton/results/TEST_OUTPUT
            volumes:
              - name: dbfolder
              - name: work
          duration: 1m 37s
          reason: Succeeded
      - name: sbom-json-check
        params:
          - name: IMAGE_URL
            value: $(tasks.build-container.results.IMAGE_URL)
          - name: IMAGE_DIGEST
            value: $(tasks.build-container.results.IMAGE_DIGEST)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: sbom-json-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/task-sbom-json-check:0.1@sha256:397cb2fb20f413dec9653134231bec86edb80806a3441081fbf473677fc40917
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'false'
            operator: in
            values:
              - 'false'
        status:
          completionTime: '2023-08-25T14:58:35Z'
          conditions:
            - lastTransitionTime: '2023-08-25T14:58:35Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: human-resources-on-pull-request-rlrj8-sbom-json-check-pod
          startTime: '2023-08-25T14:58:25Z'
          steps:
            - container: step-sbom-json-check
              imageID: >-
                quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
              name: sbom-json-check
              terminated:
                containerID: >-
                  cri-o://acef4088b52f99803e31f4fc2f934f58d60a629e4636b78b8318e5396234c11d
                exitCode: 0
                finishedAt: '2023-08-25T14:58:35Z'
                message: >-
                  [{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975515\",\"note\":\"Task
                  sbom-json-check completed: Check result for JSON check
                  result.\",\"namespace\":\"default\",\"successes\":1,\"failures\":0,\"warnings\":0}\n","type":1}]
                reason: Completed
                startedAt: '2023-08-25T14:58:29Z'
          taskResults:
            - name: TEST_OUTPUT
              type: string
              value: >
                {"result":"SUCCESS","timestamp":"1692975515","note":"Task
                sbom-json-check completed: Check result for JSON check
                result.","namespace":"default","successes":1,"failures":0,"warnings":0}
          taskSpec:
            description: >-
              Verifies the integrity and security of the Software Bill of
              Materials (SBOM) file in JSON format using CyloneDX tool.
            params:
              - description: Fully qualified image name to verify.
                name: IMAGE_URL
                type: string
              - description: Image digest.
                name: IMAGE_DIGEST
                type: string
            results:
              - description: Tekton task test output.
                name: TEST_OUTPUT
                type: string
            steps:
              - env:
                  - name: IMAGE_URL
                    value: >-
                      quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
                  - name: IMAGE_DIGEST
                    value: >-
                      sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d
                image: >-
                  quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
                name: sbom-json-check
                resources: {}
                script: >
                  #!/usr/bin/env bash

                  source /utils.sh


                  mkdir /manifests/ && cd /manifests/


                  image_with_digest="${IMAGE_URL}@${IMAGE_DIGEST}"


                  if ! oc image extract --registry-config ~/.docker/config.json
                  "${image_with_digest}" --path
                  '/root/buildinfo/content_manifests/*:/manifests/'; then
                    echo "Failed to extract manifests from image ${image_with_digest}."
                    note="Task sbom-json-check failed: Failed to extract manifests from image ${image_with_digest} with oc extract. For details, check Tekton task log."
                    ERROR_OUTPUT=$(make_result_json -r "ERROR" -t "$note")
                  fi


                  touch fail_result.txt



                  FAIL_RESULTS="$(cat fail_result.txt)"

                  if [[ -z $FAIL_RESULTS ]]

                  then
                    note="Task sbom-json-check completed: Check result for JSON check result."
                    TEST_OUTPUT=$(make_result_json -r "SUCCESS" -s 1 -t "$note")
                  else
                    echo "Failed to verify sbom-cyclonedx.json for image $IMAGE_URL with reason: $FAIL_RESULTS."
                    note="Task sbom-json-check failed: Failed to verify SBOM for image $IMAGE_URL."
                    ERROR_OUTPUT=$(make_result_json -r "FAILURE" -f 1 -t "$note")
                  fi


                  echo "${TEST_OUTPUT:-${ERROR_OUTPUT}}" | tee
                  /tekton/results/TEST_OUTPUT
                securityContext:
                  capabilities:
                    add:
                      - SETFCAP
                  runAsUser: 0
                volumeMounts:
                  - mountPath: /shared
                    name: shared
            volumes:
              - emptyDir: {}
                name: shared
          duration: 10s
          reason: Succeeded
    workspaces:
      - name: workspace
      - name: git-auth
        optional: true
  skippedTasks:
    - name: prefetch-dependencies
      reason: When Expressions evaluated to false
      whenExpressions:
        - input: 'false'
          operator: in
          values:
            - 'true'
    - name: optional-label-check
      reason: When Expressions evaluated to false
      whenExpressions:
        - input: 'true'
          operator: in
          values:
            - 'false'
    - name: sast-snyk-check
      reason: When Expressions evaluated to false
      whenExpressions:
        - input: 'false'
          operator: in
          values:
            - 'false'
        - input: ''
          operator: notin
          values:
            - ''
  startTime: '2023-08-25T14:56:28Z'
`

const trInitYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipelinesascode.tekton.dev/state: started
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipeline.tekton.dev/release: b8ad1b2
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    tekton.dev/tags: 'appstudio, hacbs'
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/c94f6df4-8137-4e16-bbe6-911a60f118d3
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    results.tekton.dev/log: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/logs/0320386e-a259-3073-aa87-256f3603bac2
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    results.tekton.dev/childReadyForDeletion: 'true'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: human-resources-on-pull-request-rlrj8-init
  uid: c94f6df4-8137-4e16-bbe6-911a60f118d3
  creationTimestamp: '2023-08-25T14:56:31Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipeline.tekton.dev/release': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:tekton.dev/tags': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:tekton.dev/pipelines.minVersion': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"4ceaef28-9295-4b0b-93fc-1a03d3feadd5"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T14:56:31Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T14:56:37Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskResults': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:steps': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T14:56:37Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
            'f:results.tekton.dev/result': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T15:05:42Z'
  namespace: rhtapuser-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: human-resources-on-pull-request-rlrj8
      uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  finalizers:
    - chains.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: started
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: human-resources
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    tekton.dev/task: init
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    tekton.dev/pipelineTask: init
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    tekton.dev/pipelineRun: human-resources-on-pull-request-rlrj8
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  params:
    - name: image-url
      value: >-
        quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    - name: rebuild
      value: 'false'
    - name: skip-checks
      value: 'false'
    - name: skip-optional
      value: 'true'
    - name: pipelinerun-name
      value: human-resources-on-pull-request-rlrj8
    - name: pipelinerun-uid
      value: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: init
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/task-init:0.1@sha256:26586a7ef08c3e86dfdaf0a5cc38dd3d70c4c02db1331b469caaed0a0f5b3d86
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
status:
  completionTime: '2023-08-25T14:56:37Z'
  conditions:
    - lastTransitionTime: '2023-08-25T14:56:37Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: human-resources-on-pull-request-rlrj8-init-pod
  startTime: '2023-08-25T14:56:31Z'
  steps:
    - container: step-init
      imageID: >-
        registry.redhat.io/openshift4/ose-cli@sha256:73df37794ffff7de1101016c23dc623e4990810390ebdabcbbfa065214352c7c
      name: init
      terminated:
        containerID: >-
          cri-o://c75af9b7fb35fad7947213d9151d0922396e426f3c8d43f92689c8349b8d1e04
        exitCode: 0
        finishedAt: '2023-08-25T14:56:37Z'
        message: >-
          [{"key":"build","value":"true","type":1},{"key":"container-registry-secret","value":"unused\n","type":1}]
        reason: Completed
        startedAt: '2023-08-25T14:56:36Z'
  taskResults:
    - name: build
      type: string
      value: 'true'
    - name: container-registry-secret
      type: string
      value: |
        unused
  taskSpec:
    description: >-
      Initialize Pipeline Task, include flags for rebuild and auth. Generates
      image repository secret used by the PipelineRun.
    params:
      - description: Image URL for build by PipelineRun
        name: image-url
        type: string
      - default: 'false'
        description: Rebuild the image if exists
        name: rebuild
        type: string
      - default: 'false'
        description: Skip checks against built image
        name: skip-checks
        type: string
      - default: 'true'
        description: 'Skip optional checks, set false if you want to run optional checks'
        name: skip-optional
        type: string
      - default: ''
        description: 'unused, should be removed in next task version'
        name: pipelinerun-name
        type: string
      - default: ''
        description: 'unused, should be removed in next task version'
        name: pipelinerun-uid
        type: string
    results:
      - description: Defines if the image in param image-url should be built
        name: build
        type: string
      - description: 'unused, should be removed in next task version'
        name: container-registry-secret
        type: string
    steps:
      - env:
          - name: IMAGE_URL
            value: >-
              quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          - name: REBUILD
            value: 'false'
          - name: SKIP_CHECKS
            value: 'false'
          - name: SKIP_OPTIONAL
            value: 'true'
        image: >-
          registry.redhat.io/openshift4/ose-cli:4.13@sha256:73df37794ffff7de1101016c23dc623e4990810390ebdabcbbfa065214352c7c
        name: init
        resources: {}
        script: >
          #!/bin/bash

          echo "Build Initialize: $IMAGE_URL"

          echo


          echo "Determine if Image Already Exists"

          # Build the image when image does not exists or rebuild is set to true

          if ! oc image info $IMAGE_URL &>/dev/null || [ "$REBUILD" == "true" ]
          || [ "$SKIP_CHECKS" == "false" ]; then
            echo -n "true" > /tekton/results/build
          else
            echo -n "false" > /tekton/results/build
          fi

          echo unused > /tekton/results/container-registry-secret
`

const trCloneYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipelinesascode.tekton.dev/state: started
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipeline.tekton.dev/release: b8ad1b2
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    tekton.dev/tags: git
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/9e8fa80f-95f7-48f2-881c-6354e637d01e
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    tekton.dev/categories: Git
    results.tekton.dev/log: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/logs/816dbb30-13cf-3c42-b629-1fa8066488a0
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    tekton.dev/platforms: 'linux/amd64,linux/s390x,linux/ppc64le,linux/arm64'
    tekton.dev/pipelines.minVersion: 0.21.0
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    results.tekton.dev/childReadyForDeletion: 'true'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    tekton.dev/displayName: git clone
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: human-resources-on-pull-request-rlrj8-clone-repository
  uid: 9e8fa80f-95f7-48f2-881c-6354e637d01e
  creationTimestamp: '2023-08-25T14:56:37Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:tekton.dev/platforms': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipeline.tekton.dev/release': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:tekton.dev/displayName': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:tekton.dev/tags': {}
            'f:tekton.dev/categories': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:tekton.dev/pipelines.minVersion': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"4ceaef28-9295-4b0b-93fc-1a03d3feadd5"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
          'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T14:56:37Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskResults': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:steps': {}
            'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T14:56:54Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T14:56:55Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T15:05:45Z'
  namespace: rhtapuser-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: human-resources-on-pull-request-rlrj8
      uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  finalizers:
    - chains.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: started
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: human-resources
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    tekton.dev/task: git-clone
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    tekton.dev/pipelineTask: clone-repository
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    tekton.dev/pipelineRun: human-resources-on-pull-request-rlrj8
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  params:
    - name: url
      value: 'https://github.com/jeff-phillips-18/human-resources'
    - name: revision
      value: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: git-clone
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/task-git-clone:0.1@sha256:1f84973a21aabea38434b1f663abc4cb2d86565a9c7aae1f90decb43a8fa48eb
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
  workspaces:
    - name: output
      persistentVolumeClaim:
        claimName: pvc-caca8b5d93
    - name: basic-auth
      secret:
        secretName: pac-gitauth-pwvj
status:
  completionTime: '2023-08-25T14:56:54Z'
  conditions:
    - lastTransitionTime: '2023-08-25T14:56:54Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: human-resources-on-pull-request-rlrj8-clone-repository-pod
  startTime: '2023-08-25T14:56:37Z'
  steps:
    - container: step-clone
      imageID: >-
        registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8@sha256:2fa0b06d52b04f377c696412e19307a9eff27383f81d87aae0b4f71672a1cd0b
      name: clone
      terminated:
        containerID: >-
          cri-o://ad6aebe8612f8776692ced04d7726d804642f3a5319004d9cd7e2b1bc780ed4b
        exitCode: 0
        finishedAt: '2023-08-25T14:56:54Z'
        message: >-
          [{"key":"commit","value":"b9651320a8c6fdd24cf35d2a1f8ea0a245235442","type":1},{"key":"url","value":"https://github.com/jeff-phillips-18/human-resources","type":1}]
        reason: Completed
        startedAt: '2023-08-25T14:56:54Z'
    - container: step-symlink-check
      imageID: >-
        registry.redhat.io/ubi9@sha256:089bd3b82a78ac45c0eed231bb58bfb43bfcd0560d9bba240fc6355502c92976
      name: symlink-check
      terminated:
        containerID: >-
          cri-o://e6a39307c093b7d07b08bcc1204e13b30e89af83a87bce7f18283ef283b078d0
        exitCode: 0
        finishedAt: '2023-08-25T14:56:54Z'
        message: >-
          [{"key":"commit","value":"b9651320a8c6fdd24cf35d2a1f8ea0a245235442","type":1},{"key":"url","value":"https://github.com/jeff-phillips-18/human-resources","type":1}]
        reason: Completed
        startedAt: '2023-08-25T14:56:54Z'
  taskResults:
    - name: commit
      type: string
      value: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    - name: url
      type: string
      value: 'https://github.com/jeff-phillips-18/human-resources'
  taskSpec:
    description: >-
      The git-clone Task will clone a repo from the provided url into the output
      Workspace. By default the repo will be cloned into the root of your
      Workspace.
    params:
      - description: Repository URL to clone from.
        name: url
        type: string
      - default: ''
        description: 'Revision to checkout. (branch, tag, sha, ref, etc...)'
        name: revision
        type: string
      - default: ''
        description: Refspec to fetch before checking out revision.
        name: refspec
        type: string
      - default: 'true'
        description: Initialize and fetch git submodules.
        name: submodules
        type: string
      - default: '1'
        description: 'Perform a shallow clone, fetching only the most recent N commits.'
        name: depth
        type: string
      - default: 'true'
        description: >-
          Set the http.sslVerify global git config. Setting this to false is
          not advised unless you are sure that you trust your git remote.
        name: sslVerify
        type: string
      - default: ''
        description: Subdirectory inside the output Workspace to clone the repo into.
        name: subdirectory
        type: string
      - default: ''
        description: >-
          Define the directory patterns to match or exclude when performing a
          sparse checkout.
        name: sparseCheckoutDirectories
        type: string
      - default: 'true'
        description: >-
          Clean out the contents of the destination directory if it already
          exists before cloning.
        name: deleteExisting
        type: string
      - default: ''
        description: HTTP proxy server for non-SSL requests.
        name: httpProxy
        type: string
      - default: ''
        description: HTTPS proxy server for SSL requests.
        name: httpsProxy
        type: string
      - default: ''
        description: Opt out of proxying HTTP/HTTPS requests.
        name: noProxy
        type: string
      - default: 'true'
        description: Log the commands that are executed during git-clone's operation.
        name: verbose
        type: string
      - default: >-
          registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8:v1.8.2-8@sha256:a538c423e7a11aae6ae582a411fdb090936458075f99af4ce5add038bb6983e8
        description: The image providing the git-init binary that this Task runs.
        name: gitInitImage
        type: string
      - default: /tekton/home
        description: >
          Absolute path to the user's home directory. Set this explicitly if you
          are running the image as a non-root user or have overridden

          the gitInitImage param with an image containing custom user
          configuration.
        name: userHome
        type: string
      - default: 'true'
        description: >
          Check symlinks in the repo. If they're pointing outside of the repo,
          the build will fail.
        name: enableSymlinkCheck
        type: string
    results:
      - description: The precise commit SHA that was fetched by this Task.
        name: commit
        type: string
      - description: The precise URL that was fetched by this Task.
        name: url
        type: string
    steps:
      - env:
          - name: HOME
            value: /tekton/home
          - name: PARAM_URL
            value: 'https://github.com/jeff-phillips-18/human-resources'
          - name: PARAM_REVISION
            value: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          - name: PARAM_REFSPEC
          - name: PARAM_SUBMODULES
            value: 'true'
          - name: PARAM_DEPTH
            value: '1'
          - name: PARAM_SSL_VERIFY
            value: 'true'
          - name: PARAM_SUBDIRECTORY
          - name: PARAM_DELETE_EXISTING
            value: 'true'
          - name: PARAM_HTTP_PROXY
          - name: PARAM_HTTPS_PROXY
          - name: PARAM_NO_PROXY
          - name: PARAM_VERBOSE
            value: 'true'
          - name: PARAM_SPARSE_CHECKOUT_DIRECTORIES
          - name: PARAM_USER_HOME
            value: /tekton/home
          - name: WORKSPACE_OUTPUT_PATH
            value: /workspace/output
          - name: WORKSPACE_SSH_DIRECTORY_BOUND
            value: 'false'
          - name: WORKSPACE_SSH_DIRECTORY_PATH
          - name: WORKSPACE_BASIC_AUTH_DIRECTORY_BOUND
            value: 'true'
          - name: WORKSPACE_BASIC_AUTH_DIRECTORY_PATH
            value: /workspace/basic-auth
        image: >-
          registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8:v1.8.2-8@sha256:a538c423e7a11aae6ae582a411fdb090936458075f99af4ce5add038bb6983e8
        name: clone
        resources: {}
        script: >
          #!/usr/bin/env sh

          set -eu


          if [ "${PARAM_VERBOSE}" = "true" ] ; then
            set -x
          fi


          if [ "${WORKSPACE_BASIC_AUTH_DIRECTORY_BOUND}" = "true" ] ; then
            if [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.git-credentials" ] && [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.gitconfig" ]; then
              cp "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.git-credentials" "${PARAM_USER_HOME}/.git-credentials"
              cp "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.gitconfig" "${PARAM_USER_HOME}/.gitconfig"
            # Compatibility with kubernetes.io/basic-auth secrets
            elif [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/username" ] && [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/password" ]; then
              HOSTNAME=$(echo $PARAM_URL | awk -F/ '{print $3}')
              echo "https://$(cat ${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/username):$(cat ${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/password)@$HOSTNAME" > "${PARAM_USER_HOME}/.git-credentials"
              echo -e "[credential \"https://$HOSTNAME\"]\n  helper = store" > "${PARAM_USER_HOME}/.gitconfig"
            else
              echo "Unknown basic-auth workspace format"
              exit 1
            fi
            chmod 400 "${PARAM_USER_HOME}/.git-credentials"
            chmod 400 "${PARAM_USER_HOME}/.gitconfig"
          fi


          if [ "${WORKSPACE_SSH_DIRECTORY_BOUND}" = "true" ] ; then
            cp -R "${WORKSPACE_SSH_DIRECTORY_PATH}" "${PARAM_USER_HOME}"/.ssh
            chmod 700 "${PARAM_USER_HOME}"/.ssh
            chmod -R 400 "${PARAM_USER_HOME}"/.ssh/*
          fi


          CHECKOUT_DIR="${WORKSPACE_OUTPUT_PATH}/${PARAM_SUBDIRECTORY}"


          cleandir() {
            # Delete any existing contents of the repo directory if it exists.
            #
            # We don't just "rm -rf ${CHECKOUT_DIR}" because ${CHECKOUT_DIR} might be "/"
            # or the root of a mounted volume.
            if [ -d "${CHECKOUT_DIR}" ] ; then
              # Delete non-hidden files and directories
              rm -rf "${CHECKOUT_DIR:?}"/*
              # Delete files and directories starting with . but excluding ..
              rm -rf "${CHECKOUT_DIR}"/.[!.]*
              # Delete files and directories starting with .. plus any other character
              rm -rf "${CHECKOUT_DIR}"/..?*
            fi
          }


          if [ "${PARAM_DELETE_EXISTING}" = "true" ] ; then
            cleandir
          fi


          test -z "${PARAM_HTTP_PROXY}" || export
          HTTP_PROXY="${PARAM_HTTP_PROXY}"

          test -z "${PARAM_HTTPS_PROXY}" || export
          HTTPS_PROXY="${PARAM_HTTPS_PROXY}"

          test -z "${PARAM_NO_PROXY}" || export NO_PROXY="${PARAM_NO_PROXY}"


          /ko-app/git-init \
            -url="${PARAM_URL}" \
            -revision="${PARAM_REVISION}" \
            -refspec="${PARAM_REFSPEC}" \
            -path="${CHECKOUT_DIR}" \
            -sslVerify="${PARAM_SSL_VERIFY}" \
            -submodules="${PARAM_SUBMODULES}" \
            -depth="${PARAM_DEPTH}" \
            -sparseCheckoutDirectories="${PARAM_SPARSE_CHECKOUT_DIRECTORIES}"
          cd "${CHECKOUT_DIR}"

          RESULT_SHA="$(git rev-parse HEAD)"

          EXIT_CODE="$?"

          if [ "${EXIT_CODE}" != 0 ] ; then
            exit "${EXIT_CODE}"
          fi

          printf "%s" "${RESULT_SHA}" > "/tekton/results/commit"

          printf "%s" "${PARAM_URL}" > "/tekton/results/url"
        securityContext:
          runAsUser: 0
      - env:
          - name: PARAM_ENABLE_SYMLINK_CHECK
            value: 'true'
          - name: PARAM_SUBDIRECTORY
          - name: WORKSPACE_OUTPUT_PATH
            value: /workspace/output
        image: 'registry.redhat.io/ubi9:9.2-696'
        name: symlink-check
        resources: {}
        script: |
          #!/usr/bin/env bash
          set -euo pipefail

          CHECKOUT_DIR="${WORKSPACE_OUTPUT_PATH}/${PARAM_SUBDIRECTORY}"
          check_symlinks() {
            FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO=false
            while read symlink
            do
              target=$(readlink -f "$symlink")
              if ! [[ "$target" =~ ^$CHECKOUT_DIR ]]; then
                echo "The cloned repository contains symlink pointing outside of the cloned repository: $symlink"
                FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO=true
              fi
            done < <(find $CHECKOUT_DIR -type l -print)
            if [ "$FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO" = true ] ; then
              return 1
            fi
          }

          if [ "${PARAM_ENABLE_SYMLINK_CHECK}" = "true" ] ; then
            echo "Running symlink check"
            check_symlinks
          fi
    workspaces:
      - description: The git repo will be cloned onto the volume backing this Workspace.
        name: output
      - description: >
          A .ssh directory with private key, known_hosts, config, etc. Copied to

          the user's home before git commands are executed. Used to authenticate

          with the git remote when performing the clone. Binding a Secret to
          this

          Workspace is strongly recommended over other volume types.
        name: ssh-directory
        optional: true
      - description: >
          A Workspace containing a .gitconfig and .git-credentials file or
          username and password.

          These will be copied to the user's home before any git commands are
          run. Any

          other files in this Workspace are ignored. It is strongly recommended

          to use ssh-directory over basic-auth whenever possible and to bind a

          Secret to this Workspace over other volume types.
        name: basic-auth
        optional: true
`

const trSbomJsonCheckYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipelinesascode.tekton.dev/state: started
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipeline.tekton.dev/release: b8ad1b2
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/6b3cae01-3953-4e97-884c-9250657339e5
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    results.tekton.dev/log: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/logs/24db4b2d-3499-3415-a625-c05741d7e878
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    results.tekton.dev/childReadyForDeletion: 'true'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: human-resources-on-pull-request-rlrj8-sbom-json-check
  uid: 6b3cae01-3953-4e97-884c-9250657339e5
  creationTimestamp: '2023-08-25T14:58:25Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipeline.tekton.dev/release': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"4ceaef28-9295-4b0b-93fc-1a03d3feadd5"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T14:58:25Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T14:58:35Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskResults': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:steps': {}
            'f:volumes': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T14:58:35Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T15:06:01Z'
  namespace: rhtapuser-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: human-resources-on-pull-request-rlrj8
      uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  finalizers:
    - chains.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: started
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: human-resources
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    tekton.dev/task: sbom-json-check
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    tekton.dev/pipelineTask: sbom-json-check
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    tekton.dev/pipelineRun: human-resources-on-pull-request-rlrj8
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  params:
    - name: IMAGE_URL
      value: >-
        quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    - name: IMAGE_DIGEST
      value: 'sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d'
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: sbom-json-check
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/task-sbom-json-check:0.1@sha256:397cb2fb20f413dec9653134231bec86edb80806a3441081fbf473677fc40917
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
status:
  completionTime: '2023-08-25T14:58:35Z'
  conditions:
    - lastTransitionTime: '2023-08-25T14:58:35Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: human-resources-on-pull-request-rlrj8-sbom-json-check-pod
  startTime: '2023-08-25T14:58:25Z'
  steps:
    - container: step-sbom-json-check
      imageID: >-
        quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
      name: sbom-json-check
      terminated:
        containerID: >-
          cri-o://acef4088b52f99803e31f4fc2f934f58d60a629e4636b78b8318e5396234c11d
        exitCode: 0
        finishedAt: '2023-08-25T14:58:35Z'
        message: >-
          [{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975515\",\"note\":\"Task
          sbom-json-check completed: Check result for JSON check
          result.\",\"namespace\":\"default\",\"successes\":1,\"failures\":0,\"warnings\":0}\n","type":1}]
        reason: Completed
        startedAt: '2023-08-25T14:58:29Z'
  taskResults:
    - name: TEST_OUTPUT
      type: string
      value: >
        {"result":"SUCCESS","timestamp":"1692975515","note":"Task
        sbom-json-check completed: Check result for JSON check
        result.","namespace":"default","successes":1,"failures":0,"warnings":0}
  taskSpec:
    description: >-
      Verifies the integrity and security of the Software Bill of Materials
      (SBOM) file in JSON format using CyloneDX tool.
    params:
      - description: Fully qualified image name to verify.
        name: IMAGE_URL
        type: string
      - description: Image digest.
        name: IMAGE_DIGEST
        type: string
    results:
      - description: Tekton task test output.
        name: TEST_OUTPUT
        type: string
    steps:
      - env:
          - name: IMAGE_URL
            value: >-
              quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          - name: IMAGE_DIGEST
            value: >-
              sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d
        image: >-
          quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
        name: sbom-json-check
        resources: {}
        script: >
          #!/usr/bin/env bash

          source /utils.sh


          mkdir /manifests/ && cd /manifests/


          image_with_digest="${IMAGE_URL}@${IMAGE_DIGEST}"


          if ! oc image extract --registry-config ~/.docker/config.json
          "${image_with_digest}" --path
          '/root/buildinfo/content_manifests/*:/manifests/'; then
            echo "Failed to extract manifests from image ${image_with_digest}."
            note="Task sbom-json-check failed: Failed to extract manifests from image ${image_with_digest} with oc extract. For details, check Tekton task log."
            ERROR_OUTPUT=$(make_result_json -r "ERROR" -t "$note")
          fi


          touch fail_result.txt

          if [ -f "sbom-cyclonedx.json" ]

          then
            result=$(echo -n $(sbom-utility validate --input-file sbom-cyclonedx.json))
            if [[ ! $result =~ "SBOM valid against JSON schema: true" ]]
            then
              echo "sbom-cyclonedx.json: $result" > fail_result.txt
            fi
          else
            echo "Cannot access sbom-cyclonedx.json: No such file or directory exists." > fail_result.txt
          fi


          FAIL_RESULTS="$(cat fail_result.txt)"

          if [[ -z $FAIL_RESULTS ]]

          then
            note="Task sbom-json-check completed: Check result for JSON check result."
            TEST_OUTPUT=$(make_result_json -r "SUCCESS" -s 1 -t "$note")
          else
            echo "Failed to verify sbom-cyclonedx.json for image $IMAGE_URL with reason: $FAIL_RESULTS."
            note="Task sbom-json-check failed: Failed to verify SBOM for image $IMAGE_URL."
            ERROR_OUTPUT=$(make_result_json -r "FAILURE" -f 1 -t "$note")
          fi


          echo "${TEST_OUTPUT:-${ERROR_OUTPUT}}" | tee
          /tekton/results/TEST_OUTPUT
        securityContext:
          capabilities:
            add:
              - SETFCAP
          runAsUser: 0
        volumeMounts:
          - mountPath: /shared
            name: shared
    volumes:
      - emptyDir: {}
        name: shared
`

const trBuildYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipelinesascode.tekton.dev/state: started
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipeline.tekton.dev/release: b8ad1b2
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    tekton.dev/tags: 'image-build, appstudio, hacbs'
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/90675680-c787-446c-bded-7f3d1469aee7
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    results.tekton.dev/log: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/logs/0c9c2f1f-d0d7-369b-99cd-0484005eecb1
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    results.tekton.dev/childReadyForDeletion: 'true'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: human-resources-on-pull-request-rlrj8-build-container
  uid: 90675680-c787-446c-bded-7f3d1469aee7
  creationTimestamp: '2023-08-25T14:56:55Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipeline.tekton.dev/release': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:tekton.dev/tags': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:tekton.dev/pipelines.minVersion': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:build.appstudio.redhat.com/build_type': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"4ceaef28-9295-4b0b-93fc-1a03d3feadd5"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
          'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T14:56:55Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskResults': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:stepTemplate':
              .: {}
              'f:env': {}
              'f:name': {}
              'f:resources': {}
            'f:steps': {}
            'f:volumes': {}
            'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T14:58:25Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T14:58:26Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T15:05:46Z'
  namespace: rhtapuser-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: human-resources-on-pull-request-rlrj8
      uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  finalizers:
    - chains.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: started
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: human-resources
    build.appstudio.redhat.com/build_type: docker
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    tekton.dev/task: buildah
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    tekton.dev/pipelineTask: build-container
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    tekton.dev/pipelineRun: human-resources-on-pull-request-rlrj8
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  params:
    - name: IMAGE
      value: >-
        quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    - name: DOCKERFILE
      value: >-
        https://raw.githubusercontent.com/devfile-samples/devfile-sample-java-springboot-basic/main/docker/Dockerfile
    - name: CONTEXT
      value: .
    - name: HERMETIC
      value: 'false'
    - name: PREFETCH_INPUT
      value: ''
    - name: IMAGE_EXPIRES_AFTER
      value: 5d
    - name: COMMIT_SHA
      value: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: buildah
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/task-buildah:0.1@sha256:e607665f13adadbd4a8d0b32768fc1b24a90884d867ecb681e15c5bc25434f71
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
  workspaces:
    - name: source
      persistentVolumeClaim:
        claimName: pvc-caca8b5d93
status:
  completionTime: '2023-08-25T14:58:25Z'
  conditions:
    - lastTransitionTime: '2023-08-25T14:58:25Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: human-resources-on-pull-request-rlrj8-build-container-pod
  startTime: '2023-08-25T14:56:55Z'
  steps:
    - container: step-build
      imageID: >-
        quay.io/redhat-appstudio/buildah@sha256:381e9bfedd59701477621da93892106873a6951b196105d3d2d85c3f6d7b569b
      name: build
      terminated:
        containerID: >-
          cri-o://2bee75648c19f53bdc6caa852fec6d8516391ac78dec27cd5edeb0b1355be656
        exitCode: 0
        finishedAt: '2023-08-25T14:58:01Z'
        reason: Completed
        startedAt: '2023-08-25T14:57:04Z'
    - container: step-sbom-syft-generate
      imageID: >-
        quay.io/redhat-appstudio/syft@sha256:244a17ce220a0b7a54c862c4fe3b72ce92799910c5eff8e94ac2f121fa5b4a53
      name: sbom-syft-generate
      terminated:
        containerID: >-
          cri-o://e4824b7c6dcd0a7aee55938e60e9e2085bbd4bbd3a6442734b4855ff6a9ec919
        exitCode: 0
        finishedAt: '2023-08-25T14:58:08Z'
        reason: Completed
        startedAt: '2023-08-25T14:58:02Z'
    - container: step-analyse-dependencies-java-sbom
      imageID: >-
        quay.io/redhat-appstudio/hacbs-jvm-build-request-processor@sha256:b198cf4b33dab59ce8ac25afd4e1001390db29ca2dec83dc8a1e21b0359ce743
      name: analyse-dependencies-java-sbom
      terminated:
        containerID: >-
          cri-o://40eac4840ea9d84672a078c389c6e872464d62614399013508708f79e35f8488
        exitCode: 0
        finishedAt: '2023-08-25T14:58:08Z'
        message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
        reason: Completed
        startedAt: '2023-08-25T14:58:08Z'
    - container: step-merge-syft-sboms
      imageID: >-
        registry.access.redhat.com/ubi9/python-39@sha256:562b4f9cd4e5abeba54c56fd2290096a71f6e10aa8c1c18f43c9d6962c2d4d41
      name: merge-syft-sboms
      terminated:
        containerID: >-
          cri-o://78e33feea2bccd7edc2a08fc6fbcae9bf3a3d6e6477aadd7b5b792b5df1c071d
        exitCode: 0
        finishedAt: '2023-08-25T14:58:09Z'
        message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
        reason: Completed
        startedAt: '2023-08-25T14:58:09Z'
    - container: step-merge-cachi2-sbom
      imageID: >-
        quay.io/redhat-appstudio/cachi2@sha256:46097f22b57e4d48a3fce96d931e08ccfe3a3e6421362d5f9353961279078eef
      name: merge-cachi2-sbom
      terminated:
        containerID: >-
          cri-o://3113046a374d117b72418467921a2f2a6026736a3f34d2c730c056c2bcd278f7
        exitCode: 0
        finishedAt: '2023-08-25T14:58:09Z'
        message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
        reason: Completed
        startedAt: '2023-08-25T14:58:09Z'
    - container: step-create-purl-sbom
      imageID: >-
        registry.access.redhat.com/ubi9/python-39@sha256:562b4f9cd4e5abeba54c56fd2290096a71f6e10aa8c1c18f43c9d6962c2d4d41
      name: create-purl-sbom
      terminated:
        containerID: >-
          cri-o://5c4cd3e48bf9a27abbeee0132e6c03fd33ca72239f52f08f00601b795b6e1c4e
        exitCode: 0
        finishedAt: '2023-08-25T14:58:09Z'
        message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
        reason: Completed
        startedAt: '2023-08-25T14:58:09Z'
    - container: step-inject-sbom-and-push
      imageID: >-
        registry.access.redhat.com/ubi9/buildah@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
      name: inject-sbom-and-push
      terminated:
        containerID: >-
          cri-o://4521e19967b852814ac76a5bd3a261dcbf1046cbdcb05091dca0aa5cc6031938
        exitCode: 0
        finishedAt: '2023-08-25T14:58:22Z'
        message: >-
          [{"key":"BASE_IMAGES_DIGESTS","value":"registry.access.redhat.com/ubi8/openjdk-17:1.15-1.1682053058@sha256:b00f687d913b8d1e027f7eabd6765de6c8d469629bef9550f10dbf207af24fe5\nregistry.access.redhat.com/ubi8/openjdk-17-runtime:1.15-1.1682053056@sha256:f921cf1f9147e4b306908f3bcb61dd215b4a51970f8db560ede02ee6a492fa99\n","type":1},{"key":"IMAGE_DIGEST","value":"sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d","type":1},{"key":"IMAGE_URL","value":"quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442","type":1},{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]
        reason: Completed
        startedAt: '2023-08-25T14:58:10Z'
    - container: step-upload-sbom
      imageID: >-
        quay.io/redhat-appstudio/cosign@sha256:c883d6f8d39148f2cea71bff4622d196d89df3e510f36c140c097b932f0dd5d5
      name: upload-sbom
      terminated:
        containerID: >-
          cri-o://bfd0cce431db96e5dd1fd3bbc69f777f20b64f72d05859054890fd2e992c6bd4
        exitCode: 0
        finishedAt: '2023-08-25T14:58:24Z'
        message: >-
          [{"key":"BASE_IMAGES_DIGESTS","value":"registry.access.redhat.com/ubi8/openjdk-17:1.15-1.1682053058@sha256:b00f687d913b8d1e027f7eabd6765de6c8d469629bef9550f10dbf207af24fe5\nregistry.access.redhat.com/ubi8/openjdk-17-runtime:1.15-1.1682053056@sha256:f921cf1f9147e4b306908f3bcb61dd215b4a51970f8db560ede02ee6a492fa99\n","type":1},{"key":"IMAGE_DIGEST","value":"sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d","type":1},{"key":"IMAGE_URL","value":"quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442","type":1},{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]
        reason: Completed
        startedAt: '2023-08-25T14:58:23Z'
  taskResults:
    - name: JAVA_COMMUNITY_DEPENDENCIES
      type: string
      value: ''
    - name: BASE_IMAGES_DIGESTS
      type: string
      value: >
        registry.access.redhat.com/ubi8/openjdk-17:1.15-1.1682053058@sha256:b00f687d913b8d1e027f7eabd6765de6c8d469629bef9550f10dbf207af24fe5

        registry.access.redhat.com/ubi8/openjdk-17-runtime:1.15-1.1682053056@sha256:f921cf1f9147e4b306908f3bcb61dd215b4a51970f8db560ede02ee6a492fa99
    - name: IMAGE_DIGEST
      type: string
      value: 'sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d'
    - name: IMAGE_URL
      type: string
      value: >-
        quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  taskSpec:
    description: >-
      Buildah task builds source code into a container image and pushes the
      image into container registry using buildah tool.

      In addition it generates a SBOM file, injects the SBOM file into final
      container image and pushes the SBOM file as separate image using cosign
      tool.

      When [Java dependency
      rebuild](https://redhat-appstudio.github.io/docs.stonesoup.io/Documentation/main/cli/proc_enabled_java_dependencies.html)
      is enabled it triggers rebuilds of Java artifacts.

      When prefetch-dependencies task was activated it is using its artifacts to
      run build in hermetic environment.
    params:
      - description: Reference of the image buildah will produce.
        name: IMAGE
        type: string
      - default: >-
          registry.access.redhat.com/ubi9/buildah:9.0.0-19@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
        description: The location of the buildah builder image.
        name: BUILDER_IMAGE
        type: string
      - default: ./Dockerfile
        description: Path to the Dockerfile to build.
        name: DOCKERFILE
        type: string
      - default: .
        description: Path to the directory to use as context.
        name: CONTEXT
        type: string
      - default: 'true'
        description: >-
          Verify the TLS on the registry endpoint (for push/pull to a non-TLS
          registry)
        name: TLSVERIFY
        type: string
      - default: ''
        description: 'unused, should be removed in next task version'
        name: DOCKER_AUTH
        type: string
      - default: 'false'
        description: Determines if build will be executed without network access.
        name: HERMETIC
        type: string
      - default: ''
        description: >-
          In case it is not empty, the prefetched content should be made
          available to the build.
        name: PREFETCH_INPUT
        type: string
      - default: ''
        description: >-
          Delete image tag after specified time. Empty means to keep the image
          tag. Time values could be something like 1h, 2d, 3w for hours, days,
          and weeks, respectively.
        name: IMAGE_EXPIRES_AFTER
        type: string
      - default: ''
        description: The image is built from this commit.
        name: COMMIT_SHA
        type: string
    results:
      - description: Digest of the image just built
        name: IMAGE_DIGEST
        type: string
      - description: Image repository where the built image was pushed
        name: IMAGE_URL
        type: string
      - description: Digests of the base images used for build
        name: BASE_IMAGES_DIGESTS
        type: string
      - description: The counting of Java components by publisher in JSON format
        name: SBOM_JAVA_COMPONENTS_COUNT
        type: string
      - description: >-
          The Java dependencies that came from community sources such as Maven
          central.
        name: JAVA_COMMUNITY_DEPENDENCIES
        type: string
    stepTemplate:
      env:
        - name: BUILDAH_FORMAT
          value: oci
        - name: STORAGE_DRIVER
          value: vfs
        - name: HERMETIC
          value: 'false'
        - name: PREFETCH_INPUT
        - name: CONTEXT
          value: .
        - name: DOCKERFILE
          value: >-
            https://raw.githubusercontent.com/devfile-samples/devfile-sample-java-springboot-basic/main/docker/Dockerfile
        - name: IMAGE
          value: >-
            quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
        - name: TLSVERIFY
          value: 'true'
        - name: IMAGE_EXPIRES_AFTER
          value: 5d
      name: ''
      resources: {}
    steps:
      - env:
          - name: COMMIT_SHA
            value: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
        image: 'quay.io/redhat-appstudio/buildah:v1.28'
        name: build
        resources:
          limits:
            cpu: '2'
            memory: 4Gi
          requests:
            cpu: 250m
            memory: 512Mi
        script: >
          if [ -e "$CONTEXT/$DOCKERFILE" ]; then
            dockerfile_path="$CONTEXT/$DOCKERFILE"
          elif [ -e "$DOCKERFILE" ]; then
            dockerfile_path="$DOCKERFILE"
          elif echo "$DOCKERFILE" | grep -q "^https\?://"; then
            echo "Fetch Dockerfile from $DOCKERFILE"
            dockerfile_path=$(mktemp --suffix=-Dockerfile)
            http_code=$(curl -s -L -w "%{http_code}" --output "$dockerfile_path" "$DOCKERFILE")
            if [ $http_code != 200 ]; then
              echo "No Dockerfile is fetched. Server responds $http_code"
              exit 1
            fi
            http_code=$(curl -s -L -w "%{http_code}" --output "$dockerfile_path.dockerignore.tmp" "$DOCKERFILE.dockerignore")
            if [ $http_code = 200 ]; then
              echo "Fetched .dockerignore from $DOCKERFILE.dockerignore"
              mv "$dockerfile_path.dockerignore.tmp" $CONTEXT/.dockerignore
            fi
          else
            echo "Cannot find Dockerfile $DOCKERFILE"
            exit 1
          fi

          if [ -n "$JVM_BUILD_WORKSPACE_ARTIFACT_CACHE_PORT_80_TCP_ADDR" ] &&
          grep -q '^\s*RUN \(./\)\?mvn' "$dockerfile_path"; then
            sed -i -e "s|^\s*RUN \(\(./\)\?mvn\(.*\)\)|RUN echo \"<settings><mirrors><mirror><id>mirror.default</id><url>http://$JVM_BUILD_WORKSPACE_ARTIFACT_CACHE_PORT_80_TCP_ADDR/v1/cache/default/0/</url><mirrorOf>*</mirrorOf></mirror></mirrors></settings>\" > /tmp/settings.yaml; \1 -s /tmp/settings.yaml|g" "$dockerfile_path"
            touch /var/lib/containers/java
          fi


          # Fixing group permission on /var/lib/containers

          chown root:root /var/lib/containers


          sed -i 's/^\s*short-name-mode\s*=\s*.*/short-name-mode = "disabled"/'
          /etc/containers/registries.conf


          # Setting new namespace to run buildah - 2^32-2

          echo 'root:1:4294967294' | tee -a /etc/subuid >> /etc/subgid


          if [ "${HERMETIC}" == "true" ]; then
            BUILDAH_ARGS="--pull=never"
            UNSHARE_ARGS="--net"
            for image in $(grep -i '^\s*FROM' "$dockerfile_path" | sed 's/--platform=\S*//' | awk '{print $2}'); do
              unshare -Ufp --keep-caps -r --map-users 1,1,65536 --map-groups 1,1,65536 -- buildah pull $image
            done
            echo "Build will be executed with network isolation"
          fi


          if [ -n "${PREFETCH_INPUT}" ]; then
            mv cachi2 /tmp/
            chmod -R go+rwX /tmp/cachi2
            VOLUME_MOUNTS="--volume /tmp/cachi2:/cachi2"
            sed -i 's|^\s*run |RUN . /cachi2/cachi2.env \&\& \\\n    |i' "$dockerfile_path"
            echo "Prefetched content will be made available"
          fi


          LABELS=(
            "--label" "build-date=$(date -u +'%Y-%m-%dT%H:%M:%S')"
            "--label" "architecture=$(uname -m)"
            "--label" "vcs-type=git"
          )

          [ -n "$COMMIT_SHA" ] && LABELS+=("--label" "vcs-ref=$COMMIT_SHA")

          [ -n "$IMAGE_EXPIRES_AFTER" ] && LABELS+=("--label"
          "quay.expires-after=$IMAGE_EXPIRES_AFTER")


          unshare -Uf $UNSHARE_ARGS --keep-caps -r --map-users 1,1,65536
          --map-groups 1,1,65536 -- buildah build \
            $VOLUME_MOUNTS \
            $BUILDAH_ARGS \
            ${LABELS[@]} \
            --tls-verify=$TLSVERIFY --no-cache \
            --ulimit nofile=4096:4096 \
            -f "$dockerfile_path" -t $IMAGE $CONTEXT

          container=$(buildah from --pull-never $IMAGE)

          buildah mount $container | tee /workspace/container_path

          echo $container > /workspace/container_name


          # Save the SBOM produced by Cachi2 so it can be merged into the final
          SBOM later

          if [ -n "${PREFETCH_INPUT}" ]; then
            cp /tmp/cachi2/output/bom.json ./sbom-cachi2.json
          fi
        securityContext:
          capabilities:
            add:
              - SETFCAP
        volumeMounts:
          - mountPath: /var/lib/containers
            name: varlibcontainers
        workingDir: /workspace/source
      - image: 'quay.io/redhat-appstudio/syft:v0.85.0'
        name: sbom-syft-generate
        resources: {}
        script: >
          syft dir:/workspace/source --file=/workspace/source/sbom-source.json
          --output=cyclonedx-json

          find $(cat /workspace/container_path) -xtype l -delete

          syft dir:$(cat /workspace/container_path)
          --file=/workspace/source/sbom-image.json --output=cyclonedx-json
        volumeMounts:
          - mountPath: /var/lib/containers
            name: varlibcontainers
      - image: >-
          quay.io/redhat-appstudio/hacbs-jvm-build-request-processor:1d417e6f1f3e68c6c537333b5759796eddae0afc
        name: analyse-dependencies-java-sbom
        resources: {}
        script: |
          if [ -f /var/lib/containers/java ]; then
            /opt/jboss/container/java/run/run-java.sh analyse-dependencies path $(cat /workspace/container_path) -s /workspace/source/sbom-image.json --task-run-name human-resources-on-pull-request-rlrj8-build-container --publishers /tekton/results/SBOM_JAVA_COMPONENTS_COUNT
            sed -i 's/^/ /' /tekton/results/SBOM_JAVA_COMPONENTS_COUNT # Workaround for SRVKP-2875
          else
            touch /tekton/results/JAVA_COMMUNITY_DEPENDENCIES
          fi
        securityContext:
          runAsUser: 0
        volumeMounts:
          - mountPath: /var/lib/containers
            name: varlibcontainers
      - image: 'registry.access.redhat.com/ubi9/python-39:1-133.1692772345'
        name: merge-syft-sboms
        resources: {}
        script: >
          #!/bin/python3

          import json


          # load SBOMs

          with open("./sbom-image.json") as f:
            image_sbom = json.load(f)

          with open("./sbom-source.json") as f:
            source_sbom = json.load(f)

          # fetch unique components from available SBOMs

          def get_identifier(component):
            return component["name"] + '@' + component.get("version", "")

          existing_components = [get_identifier(component) for component in
          image_sbom["components"]]


          for component in source_sbom["components"]:
            if get_identifier(component) not in existing_components:
              image_sbom["components"].append(component)
              existing_components.append(get_identifier(component))

          image_sbom["components"].sort(key=lambda c: get_identifier(c))


          # write the CycloneDX unified SBOM

          with open("./sbom-cyclonedx.json", "w") as f:
            json.dump(image_sbom, f, indent=4)
        securityContext:
          runAsUser: 0
        workingDir: /workspace/source
      - image: >-
          quay.io/redhat-appstudio/cachi2:0.3.0@sha256:46097f22b57e4d48a3fce96d931e08ccfe3a3e6421362d5f9353961279078eef
        name: merge-cachi2-sbom
        resources: {}
        script: |
          if [ -n "${PREFETCH_INPUT}" ]; then
            echo "Merging contents of sbom-cachi2.json into sbom-cyclonedx.json"
            /src/utils/merge_syft_sbom.py sbom-cachi2.json sbom-cyclonedx.json > sbom-temp.json
            mv sbom-temp.json sbom-cyclonedx.json
          else
            echo "Skipping step since no Cachi2 SBOM was produced"
          fi
        securityContext:
          runAsUser: 0
        workingDir: /workspace/source
      - image: 'registry.access.redhat.com/ubi9/python-39:1-133.1692772345'
        name: create-purl-sbom
        resources: {}
        script: >
          #!/bin/python3

          import json


          with open("./sbom-cyclonedx.json") as f:
            cyclonedx_sbom = json.load(f)

          purls = [{"purl": component["purl"]} for component in
          cyclonedx_sbom["components"] if "purl" in component]

          purl_content = {"image_contents": {"dependencies": purls}}


          with open("sbom-purl.json", "w") as output_file:
            json.dump(purl_content, output_file, indent=4)
        securityContext:
          runAsUser: 0
        workingDir: /workspace/source
      - image: >-
          registry.access.redhat.com/ubi9/buildah:9.0.0-19@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
        name: inject-sbom-and-push
        resources: {}
        script: >
          # Expose base image digests

          buildah images --format '{{ .Name }}:{{ .Tag }}@{{ .Digest }}' | grep
          -v $IMAGE > /tekton/results/BASE_IMAGES_DIGESTS


          base_image_name=$(buildah inspect --format '{{ index .ImageAnnotations
          "org.opencontainers.image.base.name"}}' $IMAGE | cut -f1 -d'@')

          base_image_digest=$(buildah inspect --format '{{ index
          .ImageAnnotations "org.opencontainers.image.base.digest"}}' $IMAGE)

          container=$(buildah from --pull-never $IMAGE)

          buildah copy $container sbom-cyclonedx.json sbom-purl.json
          /root/buildinfo/content_manifests/

          buildah config -a
          org.opencontainers.image.base.name=${base_image_name} -a
          org.opencontainers.image.base.digest=${base_image_digest} $container

          buildah commit $container $IMAGE


          status=-1

          max_run=5

          sleep_sec=10

          for run in $(seq 1 $max_run); do
            status=0
            [ "$run" -gt 1 ] && sleep $sleep_sec
            echo "Pushing sbom image to registry"
            buildah push \
              --tls-verify=$TLSVERIFY \
              --digestfile /workspace/source/image-digest $IMAGE \
              docker://$IMAGE && break || status=$?
          done

          if [ "$status" -ne 0 ]; then
              echo "Failed to push sbom image to registry after ${max_run} tries"
              exit 1
          fi


          cat "/workspace/source"/image-digest | tee
          /tekton/results/IMAGE_DIGEST

          echo -n "$IMAGE" | tee /tekton/results/IMAGE_URL
        securityContext:
          capabilities:
            add:
              - SETFCAP
          runAsUser: 0
        volumeMounts:
          - mountPath: /var/lib/containers
            name: varlibcontainers
        workingDir: /workspace/source
      - args:
          - attach
          - sbom
          - '--sbom'
          - sbom-cyclonedx.json
          - '--type'
          - cyclonedx
          - >-
            quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
        image: 'quay.io/redhat-appstudio/cosign:v2.1.1'
        name: upload-sbom
        resources: {}
        workingDir: /workspace/source
    volumes:
      - emptyDir: {}
        name: varlibcontainers
    workspaces:
      - description: Workspace containing the source code to build.
        name: source
`

const trDeprecatedBaseImgCheck = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipelinesascode.tekton.dev/state: started
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipeline.tekton.dev/release: b8ad1b2
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    tekton.dev/tags: 'appstudio, hacbs'
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/6a6b6634-ff51-4f4d-82ea-e641190fa6c9
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    results.tekton.dev/log: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/logs/70f2ff23-1705-304c-b61f-3413fe4024b7
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    results.tekton.dev/childReadyForDeletion: 'true'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: humeaa4363f279f84d4e98a2fee1964cf0b-deprecated-base-image-check
  uid: 6a6b6634-ff51-4f4d-82ea-e641190fa6c9
  creationTimestamp: '2023-08-25T14:58:25Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipeline.tekton.dev/release': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:tekton.dev/tags': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:tekton.dev/pipelines.minVersion': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"4ceaef28-9295-4b0b-93fc-1a03d3feadd5"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
          'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T14:58:25Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskResults': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:steps': {}
            'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T14:58:48Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T14:58:49Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T15:05:51Z'
  namespace: rhtapuser-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: human-resources-on-pull-request-rlrj8
      uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  finalizers:
    - chains.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: started
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: human-resources
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    tekton.dev/task: deprecated-image-check
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    tekton.dev/pipelineTask: deprecated-base-image-check
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    tekton.dev/pipelineRun: human-resources-on-pull-request-rlrj8
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  params:
    - name: BASE_IMAGES_DIGESTS
      value: >
        registry.access.redhat.com/ubi8/openjdk-17:1.15-1.1682053058@sha256:b00f687d913b8d1e027f7eabd6765de6c8d469629bef9550f10dbf207af24fe5

        registry.access.redhat.com/ubi8/openjdk-17-runtime:1.15-1.1682053056@sha256:f921cf1f9147e4b306908f3bcb61dd215b4a51970f8db560ede02ee6a492fa99
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: deprecated-image-check
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/task-deprecated-image-check:0.2@sha256:58d16de95b4ca597f7f860fb85d6206e549910fa7a8d2a2cc229558f791ad329
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
  workspaces:
    - name: test-ws
      persistentVolumeClaim:
        claimName: pvc-caca8b5d93
status:
  completionTime: '2023-08-25T14:58:48Z'
  conditions:
    - lastTransitionTime: '2023-08-25T14:58:48Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: humeaa4363f279f84d4e98a2feefb857604072fb4347373bf5b615f86ad-pod
  startTime: '2023-08-25T14:58:25Z'
  steps:
    - container: step-query-pyxis
      imageID: >-
        registry.access.redhat.com/ubi8/ubi-minimal@sha256:7394c071ed74ace08cfd51f881c94067fa7a570e7f7e4a0ef0aff1b4f6a2a949
      name: query-pyxis
      terminated:
        containerID: >-
          cri-o://521ca0e70485705fdf7b7b7b15e517019a6099ebb6099f41128d980f054ea3e6
        exitCode: 0
        finishedAt: '2023-08-25T14:58:47Z'
        message: >-
          [{"key":"PYXIS_HTTP_CODE","value":"200 registry.access.redhat.com
          ubi8/openjdk-17\n200 registry.access.redhat.com
          ubi8/openjdk-17-runtime\n","type":1}]
        reason: Completed
        startedAt: '2023-08-25T14:58:47Z'
    - container: step-run-conftest
      imageID: >-
        quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
      name: run-conftest
      terminated:
        containerID: >-
          cri-o://8f8de37355b5e0fa6d6fbc0f55ed28dc2cd9ad600dc6cc27a62a45434fb4ecb3
        exitCode: 0
        finishedAt: '2023-08-25T14:58:48Z'
        message: >-
          [{"key":"PYXIS_HTTP_CODE","value":"200 registry.access.redhat.com
          ubi8/openjdk-17\n200 registry.access.redhat.com
          ubi8/openjdk-17-runtime\n","type":1},{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975528\",\"note\":\"Task
          deprecated-image-check completed: Check result for task
          result.\",\"namespace\":\"required_checks\",\"successes\":2,\"failures\":0,\"warnings\":0}\n","type":1}]
        reason: Completed
        startedAt: '2023-08-25T14:58:48Z'
  taskResults:
    - name: PYXIS_HTTP_CODE
      type: string
      value: |
        200 registry.access.redhat.com ubi8/openjdk-17
        200 registry.access.redhat.com ubi8/openjdk-17-runtime
    - name: TEST_OUTPUT
      type: string
      value: >
        {"result":"SUCCESS","timestamp":"1692975528","note":"Task
        deprecated-image-check completed: Check result for task
        result.","namespace":"required_checks","successes":2,"failures":0,"warnings":0}
  taskSpec:
    description: >-
      Identifies the unmaintained and potentially insecure deprecated base
      images. Pyxis API collects metadata from image repository, and Conftest
      applies supplied policy to identify the deprecated images using that
      metadata.
    params:
      - default: /project/repository/
        description: Path to directory containing Conftest policies.
        name: POLICY_DIR
        type: string
      - default: required_checks
        description: Namespace for Conftest policy.
        name: POLICY_NAMESPACE
        type: string
      - description: Digests of base build images.
        name: BASE_IMAGES_DIGESTS
        type: string
    results:
      - description: HTTP code returned by Pyxis API endpoint.
        name: PYXIS_HTTP_CODE
        type: string
      - description: Tekton task test output.
        name: TEST_OUTPUT
        type: string
    steps:
      - env:
          - name: BASE_IMAGES_DIGESTS
            value: >
              registry.access.redhat.com/ubi8/openjdk-17:1.15-1.1682053058@sha256:b00f687d913b8d1e027f7eabd6765de6c8d469629bef9550f10dbf207af24fe5

              registry.access.redhat.com/ubi8/openjdk-17-runtime:1.15-1.1682053056@sha256:f921cf1f9147e4b306908f3bcb61dd215b4a51970f8db560ede02ee6a492fa99
        image: >-
          registry.access.redhat.com/ubi8/ubi-minimal:8.8-1037@sha256:8d43664c250c72d35af8498c7ff76a9f0d42f16b9b3b29f0caa747121778de0e
        name: query-pyxis
        resources: {}
        script: >
          #!/usr/bin/env bash

          readarray -t IMAGE_ARRAY < <(echo -n "$BASE_IMAGES_DIGESTS" | sed
          's/\\n/\'$'\n''/g')

          for BASE_IMAGE in ${IMAGE_ARRAY[@]};

          do
            IFS=:'/' read -r IMAGE_REGISTRY IMAGE_WITH_TAG <<< $BASE_IMAGE; echo "[$IMAGE_REGISTRY] [$IMAGE_WITH_TAG]"
            IMAGE_REPOSITORY=echo $IMAGE_WITH_TAG | cut -d ":" -f1
            IMAGE_REGISTRY=${IMAGE_REGISTRY//registry.redhat.io/registry.access.redhat.com}
            export IMAGE_REPO_PATH=/workspace/test-ws/${IMAGE_REPOSITORY}
            mkdir -p ${IMAGE_REPO_PATH}
            echo "Querying Pyxis for $BASE_IMAGE."
            http_code=$(curl -s -k -o ${IMAGE_REPO_PATH}/repository_data.json -w '%{http_code}' "https://catalog.redhat.com/api/containers/v1/repositories/registry/${IMAGE_REGISTRY}/repository/${IMAGE_REPOSITORY}")
            echo "Response code: $http_code."
            echo $http_code $IMAGE_REGISTRY $IMAGE_REPOSITORY>> /tekton/results/PYXIS_HTTP_CODE
          done
      - env:
          - name: POLICY_DIR
            value: /project/repository/
          - name: POLICY_NAMESPACE
            value: required_checks
        image: >-
          quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
        name: run-conftest
        resources: {}
        script: >
          #!/usr/bin/env sh

          source /utils.sh


          success_counter=0

          failure_counter=0

          error_counter=0

          if [ ! -f /tekton/results/PYXIS_HTTP_CODE ]; then
            error_counter=$((error_counter++))
          fi

          while IFS= read -r line

          do
            IFS=:' ' read -r http_code IMAGE_REGISTRY IMAGE_REPOSITORY <<< $line; echo "[$http_code] [$IMAGE_REGISTRY] [$IMAGE_REPOSITORY]"
            export IMAGE_REPO_PATH=/workspace/test-ws/${IMAGE_REPOSITORY}
            if [ "$http_code" == "200" ];
            then
              echo "Running conftest using $POLICY_DIR policy, $POLICY_NAMESPACE namespace."
              /usr/bin/conftest test --no-fail ${IMAGE_REPO_PATH}/repository_data.json \
              --policy $POLICY_DIR --namespace $POLICY_NAMESPACE \
              --output=json 2> ${IMAGE_REPO_PATH}/stderr.txt | tee ${IMAGE_REPO_PATH}/deprecated_image_check_output.json

              failure_counter=$((failure_counter+$(jq -r '.[].failures|length' ${IMAGE_REPO_PATH}/deprecated_image_check_output.json)))
              success_counter=$((success_counter+$(jq -r '.[].successes' ${IMAGE_REPO_PATH}/deprecated_image_check_output.json)))

            elif [ "$http_code" == "404" ];
            then
              echo "Registry/image ${IMAGE_REGISTRY}/${IMAGE_REPOSITORY} not found in Pyxis." >> /workspace/test-ws/stderr.txt
              cat /workspace/test-ws/stderr.txt
            else
              echo "Unexpected error HTTP code $http_code) occurred for registry/image ${IMAGE_REGISTRY}/${IMAGE_REPOSITORY}." >> /workspace/test-ws/stderr.txt
              cat /workspace/test-ws/stderr.txt
              error_counter=$((error_counter++))
              exit 0
            fi
          done < /tekton/results/PYXIS_HTTP_CODE


          note="Task deprecated-image-check failed: Command conftest failed. For
          details, check Tekton task log."

          ERROR_OUTPUT=$(make_result_json -r ERROR -n "$POLICY_NAMESPACE" -t
          "$note")

          if [[ "$error_counter" == 0 && "$success_counter" > 0 ]];

          then
            if [[ "${failure_counter}" -gt 0 ]]; then RES="FAILURE"; else RES="SUCCESS"; fi
            note="Task deprecated-image-check completed: Check result for task result."
            TEST_OUTPUT=$(make_result_json \
              -r "${RES}" -n "$POLICY_NAMESPACE" \
              -s "${success_counter}" -f "${failure_counter}" -t "$note")
          fi

          echo "${TEST_OUTPUT:-${ERROR_OUTPUT}}" | tee
          /tekton/results/TEST_OUTPUT
    workspaces:
      - name: test-ws
`

const trInspectImgYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipelinesascode.tekton.dev/state: started
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipeline.tekton.dev/release: b8ad1b2
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    tekton.dev/tags: 'appstudio, hacbs'
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/72aa566d-731e-4b73-93cc-ca45e4ce7e9b
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    results.tekton.dev/log: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/logs/f90ff6b0-c7b2-3261-8c41-fb9fa6e3d4a7
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    results.tekton.dev/childReadyForDeletion: 'true'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: human-resources-on-pull-request-rlrj8-inspect-image
  uid: 72aa566d-731e-4b73-93cc-ca45e4ce7e9b
  creationTimestamp: '2023-08-25T14:58:25Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipeline.tekton.dev/release': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:tekton.dev/tags': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:tekton.dev/pipelines.minVersion': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"4ceaef28-9295-4b0b-93fc-1a03d3feadd5"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
          'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T14:58:25Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskResults': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:steps': {}
            'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T14:58:48Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T14:58:49Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T15:05:49Z'
  namespace: rhtapuser-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: human-resources-on-pull-request-rlrj8
      uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  finalizers:
    - chains.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: started
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: human-resources
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    tekton.dev/task: inspect-image
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    tekton.dev/pipelineTask: inspect-image
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    tekton.dev/pipelineRun: human-resources-on-pull-request-rlrj8
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  params:
    - name: IMAGE_URL
      value: >-
        quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    - name: IMAGE_DIGEST
      value: 'sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d'
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: inspect-image
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/task-inspect-image:0.1@sha256:bbc286f0a2ad94e671ceb9d0f1debd96f36b8c38c1147c5030957820b4125fc6
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
  workspaces:
    - name: source
      persistentVolumeClaim:
        claimName: pvc-caca8b5d93
status:
  completionTime: '2023-08-25T14:58:48Z'
  conditions:
    - lastTransitionTime: '2023-08-25T14:58:48Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: human-resources-on-pull-request-rlrj8-inspect-image-pod
  startTime: '2023-08-25T14:58:25Z'
  steps:
    - container: step-inspect-image
      imageID: >-
        quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
      name: inspect-image
      terminated:
        containerID: >-
          cri-o://8d1abe9b4cb1cda026b9da353d86abbdce8fc96cbc1029c7d1b601ed218f8f01
        exitCode: 0
        finishedAt: '2023-08-25T14:58:48Z'
        message: >-
          [{"key":"BASE_IMAGE","value":"registry.access.redhat.com/ubi8/openjdk-17-runtime@sha256:14de89e89efc97aee3b50141108b7833708c3a93ad90bf89940025ab5267ba86","type":1},{"key":"BASE_IMAGE_REPOSITORY","value":"ubi8/openjdk-17-runtime","type":1},{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975528\",\"note\":\"Task
          inspect-image completed: Check inspected JSON files under
          /workspace/source/hacbs/inspect-image.\",\"namespace\":\"default\",\"successes\":1,\"failures\":0,\"warnings\":0}\n","type":1}]
        reason: Completed
        startedAt: '2023-08-25T14:58:47Z'
  taskResults:
    - name: BASE_IMAGE
      type: string
      value: >-
        registry.access.redhat.com/ubi8/openjdk-17-runtime@sha256:14de89e89efc97aee3b50141108b7833708c3a93ad90bf89940025ab5267ba86
    - name: BASE_IMAGE_REPOSITORY
      type: string
      value: ubi8/openjdk-17-runtime
    - name: TEST_OUTPUT
      type: string
      value: >
        {"result":"SUCCESS","timestamp":"1692975528","note":"Task inspect-image
        completed: Check inspected JSON files under
        /workspace/source/hacbs/inspect-image.","namespace":"default","successes":1,"failures":0,"warnings":0}
  taskSpec:
    description: >-
      Inspects and analyzes manifest data of the container's source image, and
      its base image (if available) using Skopeo. An image's manifest data
      contains information about the layers that make up the image, the
      platforms for which the image is intended, and other metadata about the
      image.
    params:
      - description: Fully qualified image name.
        name: IMAGE_URL
        type: string
      - description: Image digest.
        name: IMAGE_DIGEST
        type: string
      - default: ''
        description: 'unused, should be removed in next task version'
        name: DOCKER_AUTH
        type: string
    results:
      - description: Base image source image is built from.
        name: BASE_IMAGE
        type: string
      - description: Base image repository URL.
        name: BASE_IMAGE_REPOSITORY
        type: string
      - description: Tekton task test output.
        name: TEST_OUTPUT
        type: string
    steps:
      - env:
          - name: IMAGE_URL
            value: >-
              quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          - name: IMAGE_DIGEST
            value: >-
              sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d
        image: >-
          quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
        name: inspect-image
        resources: {}
        script: >
          #!/usr/bin/env bash

          source /utils.sh

          IMAGE_INSPECT=image_inspect.json

          BASE_IMAGE_INSPECT=base_image_inspect.json

          RAW_IMAGE_INSPECT=raw_image_inspect.json


          IMAGE_URL="${IMAGE_URL}@${IMAGE_DIGEST}"

          # Given a tag and a the digest in the IMAGE_URL we opt to use the
          digest alone

          # this is because containers/image currently doesn't support image
          references

          # that contain both. See
          https://github.com/containers/image/issues/1736

          if [[ "${IMAGE_URL}" == *":"*"@"* ]]; then
            IMAGE_URL="${IMAGE_URL/:*@/@}"
          fi


          status=-1

          max_run=5

          sleep_sec=10

          for run in $(seq 1 $max_run); do
            status=0
            [ "$run" -gt 1 ] && sleep $sleep_sec  # skip last sleep
            echo "Inspecting manifest for source image ${IMAGE_URL} (try $run/$max_run)."
            skopeo inspect --no-tags docker://"${IMAGE_URL}" > $IMAGE_INSPECT && break || status=$?
          done

          if [ "$status" -ne 0 ]; then
              echo "Failed to inspect image ${IMAGE_URL}"
              note="Task inspect-image failed: Encountered errors while inspecting image. For details, check Tekton task log."
              TEST_OUTPUT=$(make_result_json -r ERROR -t "$note")
              echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
              exit 0
          fi

          echo "Image ${IMAGE_URL} metadata:"

          cat "$IMAGE_INSPECT"


          for run in $(seq 1 $max_run); do
            status=0
            [ "$run" -gt 1 ] && sleep $sleep_sec  # skip last sleep
            echo "Inspecting raw image manifest ${IMAGE_URL} (try $run/$max_run)."
            skopeo inspect --no-tags --raw docker://"${IMAGE_URL}" > $RAW_IMAGE_INSPECT && break || status=$?
          done

          if [ "$status" -ne 0 ]; then
              echo "Failed to get raw metadata of image ${IMAGE_URL}"
              note="Task inspect-image failed: Encountered errors while inspecting image. For details, check Tekton task log."
              TEST_OUTPUT=$(make_result_json -r ERROR -t "$note")
              echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
              exit 0
          fi

          echo "Image ${IMAGE_URL} raw metadata:"

          cat "$RAW_IMAGE_INSPECT" | jq  # jq for readable formatting


          echo "Getting base image manifest for source image ${IMAGE_URL}."

          BASE_IMAGE_NAME="$(jq -r
          ".annotations.\"org.opencontainers.image.base.name\""
          $RAW_IMAGE_INSPECT)"

          BASE_IMAGE_DIGEST="$(jq -r
          ".annotations.\"org.opencontainers.image.base.digest\""
          $RAW_IMAGE_INSPECT)"

          if [ $BASE_IMAGE_NAME == 'null' ]; then
            echo "Cannot get base image info from annotations."
            BASE_IMAGE_NAME="$(jq -r ".Labels.\"org.opencontainers.image.base.name\"" $IMAGE_INSPECT)"
            BASE_IMAGE_DIGEST="$(jq -r ".annotations.\"org.opencontainers.image.base.digest\"" $IMAGE_INSPECT)"
            if [ "$BASE_IMAGE_NAME" == 'null' ]; then
              echo "Cannot get base image info from Labels. For details, check source image ${IMAGE_URL}."
              exit 0
            fi
          fi

          if [ -z "$BASE_IMAGE_NAME" ]; then
            echo "Source image ${IMAGE_URL} is built from scratch, so there is no base image."
            exit 0
          fi


          BASE_IMAGE="${BASE_IMAGE_NAME%:*}@$BASE_IMAGE_DIGEST"

          echo "Detected base image: $BASE_IMAGE"

          echo -n "$BASE_IMAGE" > /tekton/results/BASE_IMAGE


          for run in $(seq 1 $max_run); do
            status=0
            [ "$run" -gt 1 ] && sleep $sleep_sec  # skip last sleep
            echo "Inspecting base image ${BASE_IMAGE} (try $run/$max_run)."
            skopeo inspect --no-tags "docker://$BASE_IMAGE"  > $BASE_IMAGE_INSPECT && break || status=$?
          done

          if [ "$status" -ne 0 ]; then
              echo "Failed to inspect base image ${BASE_IMAGE}"
              note="Task inspect-image failed: Encountered errors while inspecting image. For details, check Tekton task log."
              TEST_OUTPUT=$(make_result_json -r ERROR -t "$note")
              echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
              exit 0
          fi


          BASE_IMAGE_REPOSITORY="$(jq -r '.Name | sub("[^/]+/"; "") |
          sub("[:@].*"; "")' "$BASE_IMAGE_INSPECT")"

          echo "Detected base image repository: $BASE_IMAGE_REPOSITORY"

          echo -n "$BASE_IMAGE_REPOSITORY" >
          /tekton/results/BASE_IMAGE_REPOSITORY


          note="Task inspect-image completed: Check inspected JSON files under
          /workspace/source/hacbs/inspect-image."

          TEST_OUTPUT=$(make_result_json -r SUCCESS -s 1 -t "$note")

          echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
        securityContext:
          capabilities:
            add:
              - SETFCAP
          runAsUser: 0
        workingDir: /workspace/source/hacbs/inspect-image
    workspaces:
      - name: source
`

const trLabelYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipelinesascode.tekton.dev/state: started
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipeline.tekton.dev/release: b8ad1b2
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    tekton.dev/tags: 'appstudio, hacbs'
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/d0a62c6a-2b17-4c75-900e-e8ec99585253
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    results.tekton.dev/log: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/logs/7e5c6e5b-8437-37bc-9937-11216adef86e
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    results.tekton.dev/childReadyForDeletion: 'true'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: human-resources-on-pull-request-rlrj8-label-check
  uid: d0a62c6a-2b17-4c75-900e-e8ec99585253
  creationTimestamp: '2023-08-25T14:58:49Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipeline.tekton.dev/release': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:tekton.dev/tags': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:tekton.dev/pipelines.minVersion': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"4ceaef28-9295-4b0b-93fc-1a03d3feadd5"}': {}
        'f:spec':
          .: {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
          'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T14:58:49Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T14:59:15Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskResults': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:steps': {}
            'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T14:59:15Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T15:06:03Z'
  namespace: rhtapuser-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: human-resources-on-pull-request-rlrj8
      uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  finalizers:
    - chains.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: started
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: human-resources
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    tekton.dev/task: label-check
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    tekton.dev/pipelineTask: label-check
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    tekton.dev/pipelineRun: human-resources-on-pull-request-rlrj8
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: label-check
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/task-label-check:0.1@sha256:0c0739fdda24cd1e3587bbab9b07d4493efc21884baac7723f4b446e95bf1fd3
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
  workspaces:
    - name: workspace
      persistentVolumeClaim:
        claimName: pvc-caca8b5d93
status:
  completionTime: '2023-08-25T14:59:15Z'
  conditions:
    - lastTransitionTime: '2023-08-25T14:59:15Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: human-resources-on-pull-request-rlrj8-label-check-pod
  startTime: '2023-08-25T14:58:49Z'
  steps:
    - container: step-surface-level-checks-required-labels
      imageID: >-
        quay.io/redhat-appstudio/hacbs-test@sha256:82b43bffe4eacc717239424f64478b18f36528df47c2d11df3a8d031e81a3c67
      name: surface-level-checks-required-labels
      terminated:
        containerID: >-
          cri-o://446c06967a1115140d01c2c34bb74cdf00ba5ca6af970d8891c796f9d4b762b6
        exitCode: 0
        finishedAt: '2023-08-25T14:59:14Z'
        message: >-
          [{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975554\",\"note\":\"For
          details, check Tekton task
          log.\",\"namespace\":\"required_checks\",\"successes\":21,\"failures\":0,\"warnings\":0}\n","type":1}]
        reason: Completed
        startedAt: '2023-08-25T14:59:14Z'
  taskResults:
    - name: TEST_OUTPUT
      type: string
      value: >
        {"result":"SUCCESS","timestamp":"1692975554","note":"For details, check
        Tekton task
        log.","namespace":"required_checks","successes":21,"failures":0,"warnings":0}
  taskSpec:
    description: >-
      Verifies whether an image contains the best practice labels using
      Conftest.
    params:
      - default: /project/image/
        description: Path to directory containing Conftest policies.
        name: POLICY_DIR
        type: string
      - default: required_checks
        description: Namespace for Conftest policy.
        name: POLICY_NAMESPACE
        type: string
    results:
      - description: Tekton task test output.
        name: TEST_OUTPUT
        type: string
    steps:
      - env:
          - name: POLICY_NAMESPACE
            value: required_checks
          - name: POLICY_DIR
            value: /project/image/
        image: >-
          quay.io/redhat-appstudio/hacbs-test:v1.1.0@sha256:82b43bffe4eacc717239424f64478b18f36528df47c2d11df3a8d031e81a3c67
        name: surface-level-checks-required-labels
        resources: {}
        script: >
          #!/usr/bin/env bash


          . /utils.sh

          if [ ! -s ../inspect-image/image_inspect.json ]; then
            echo "File $(workspaces.source.path)/hacbs/inspect-image/image_inspect.json did not generate correctly. Check task inspect-image log."
            note="Task label-check failed: $(workspaces.source.path)/hacbs/inspect-image/image_inspect.json did not generate correctly. For details, check Tekton task result TEST_OUTPUT in task inspect-image."
            TEST_OUTPUT=$(make_result_json -r ERROR -t "$note")
            echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
            exit 0
          fi


          CONFTEST_OPTIONS=""

          if [ -s "../inspect-image/base_image_inspect.json" ]; then
            CONFTEST_OPTIONS="-d=../inspect-image/base_image_inspect.json"
          fi


          echo "Running conftest using $POLICY_DIR policy, $POLICY_NAMESPACE
          namespace."

          /usr/bin/conftest test --no-fail ../inspect-image/image_inspect.json
          "${CONFTEST_OPTIONS}" \

          --policy $POLICY_DIR --namespace $POLICY_NAMESPACE \

          --output=json 2> stderr.txt | tee label_check_output.json


          if [ ! -z $(cat stderr.txt) ]; then
            echo "label-check test encountered the following error:"
            cat stderr.txt
            note="Task label-check failed: Command conftest failed. For details, check Tekton task log."
            ERROR_OUTPUT=$(make_result_json -r "ERROR" -t "$note")
          fi


          TEST_OUTPUT=

          parse_test_output label-check conftest label_check_output.json || true


          echo "${TEST_OUTPUT:-${ERROR_OUTPUT}}" | tee
          /tekton/results/TEST_OUTPUT
        securityContext:
          capabilities:
            add:
              - SETFCAP
        workingDir: /workspace/workspace/hacbs/label-check-required_checks
    workspaces:
      - name: workspace
`

const trClamavYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipelinesascode.tekton.dev/state: started
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipeline.tekton.dev/release: b8ad1b2
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    tekton.dev/tags: 'virus, appstudio, hacbs'
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/27e8ab32-f925-4e7e-bc76-ff1c535d95a3
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    results.tekton.dev/log: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/logs/4ce8fc10-b5a4-388e-beab-d234b2ab9da8
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    results.tekton.dev/childReadyForDeletion: 'true'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: human-resources-on-pull-request-rlrj8-clamav-scan
  uid: 27e8ab32-f925-4e7e-bc76-ff1c535d95a3
  creationTimestamp: '2023-08-25T14:58:25Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipeline.tekton.dev/release': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:tekton.dev/tags': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:tekton.dev/pipelines.minVersion': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"4ceaef28-9295-4b0b-93fc-1a03d3feadd5"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T14:58:25Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T15:01:02Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          'f:sidecars': {}
          'f:conditions': {}
          .: {}
          'f:steps': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:sidecars': {}
            'f:steps': {}
            'f:volumes': {}
          'f:startTime': {}
          'f:podName': {}
          'f:taskResults': {}
          'f:completionTime': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T15:01:02Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T15:05:57Z'
  namespace: rhtapuser-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: human-resources-on-pull-request-rlrj8
      uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  finalizers:
    - chains.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: started
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: human-resources
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    tekton.dev/task: clamav-scan
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    tekton.dev/pipelineTask: clamav-scan
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    tekton.dev/pipelineRun: human-resources-on-pull-request-rlrj8
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  params:
    - name: image-digest
      value: 'sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d'
    - name: image-url
      value: >-
        quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: clamav-scan
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/task-clamav-scan:0.1@sha256:cd4e301dd849cbdf7b8e38fd8f4915970b5b60174770df632a6b38ea93028d44
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
status:
  completionTime: '2023-08-25T15:01:02Z'
  conditions:
    - lastTransitionTime: '2023-08-25T15:01:02Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: human-resources-on-pull-request-rlrj8-clamav-scan-pod
  sidecars:
    - container: sidecar-database
      imageID: >-
        quay.io/redhat-appstudio/clamav-db@sha256:703c928d5d34a6004f50e28301d4aa642d3eb18edaa6a697dc73fbe72c46ffe5
      name: database
      terminated:
        containerID: >-
          cri-o://9105b69f157a8f5a92856ab8dd27c5d392d32ab7ad00c7442825f3112b20f04f
        exitCode: 0
        finishedAt: '2023-08-25T14:59:29Z'
        reason: Completed
        startedAt: '2023-08-25T14:59:29Z'
  startTime: '2023-08-25T14:59:25Z'
  steps:
    - container: step-extract-and-scan-image
      imageID: >-
        quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
      name: extract-and-scan-image
      terminated:
        containerID: >-
          cri-o://16fa2bfc99dc510afdc9983de8501b57013101c128ba467905d118134ef80a0e
        exitCode: 0
        finishedAt: '2023-08-25T15:01:01Z'
        reason: Completed
        startedAt: '2023-08-25T14:59:31Z'
    - container: step-modify-clam-output-to-json
      imageID: >-
        quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
      name: modify-clam-output-to-json
      terminated:
        containerID: >-
          cri-o://0dafda6e48f6d18bb0aa3421d0dc64804eea9c35fb2f2eeaf4d4e09f01ac83ea
        exitCode: 0
        finishedAt: '2023-08-25T15:01:02Z'
        reason: Completed
        startedAt: '2023-08-25T15:01:01Z'
    - container: step-store-hacbs-test-output-result
      imageID: >-
        quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
      name: store-hacbs-test-output-result
      terminated:
        containerID: >-
          cri-o://24ae0f28387e6b1287857396af6be637a3c03a308e5b3d1396451630e1c1514c
        exitCode: 0
        finishedAt: '2023-08-25T15:01:02Z'
        message: >-
          [{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975662\",\"note\":\"Task
          clamav-scan completed: Check result for antivirus scan
          result.\",\"namespace\":\"default\",\"successes\":1,\"failures\":0,\"warnings\":0}\n","type":1}]
        reason: Completed
        startedAt: '2023-08-25T15:01:02Z'
  taskResults:
    - name: TEST_OUTPUT
      type: string
      value: >
        {"result":"SUCCESS","timestamp":"1692975662","note":"Task clamav-scan
        completed: Check result for antivirus scan
        result.","namespace":"default","successes":1,"failures":0,"warnings":0}
  taskSpec:
    description: >-
      Scans the content of container images for viruses, malware, and other
      malicious content using ClamAV antivirus scanner.
    params:
      - description: Image digest to scan.
        name: image-digest
        type: string
      - description: Image URL.
        name: image-url
        type: string
      - default: ''
        description: unused
        name: docker-auth
        type: string
    results:
      - description: Tekton task test output.
        name: TEST_OUTPUT
        type: string
    sidecars:
      - image: 'quay.io/redhat-appstudio/clamav-db:v1'
        imagePullPolicy: Always
        name: database
        resources: {}
        script: |
          #!/usr/bin/env bash
          clamscan --version
          cp -r /var/lib/clamav/* /tmp/clamdb
        volumeMounts:
          - mountPath: /tmp/clamdb
            name: dbfolder
    steps:
      - env:
          - name: HOME
            value: /work
          - name: IMAGE_URL
            value: >-
              quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          - name: IMAGE_DIGEST
            value: >-
              sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d
        image: >-
          quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
        name: extract-and-scan-image
        resources:
          limits:
            cpu: '2'
            memory: 4Gi
          requests:
            cpu: 10m
            memory: 512Mi
        script: >
          imagewithouttag=$(echo $IMAGE_URL | sed "s/\(.*\):.*/\1/" | tr -d
          '\n')


          # strip new-line escape symbol from parameter and save it to variable

          imageanddigest=$(echo $imagewithouttag@$IMAGE_DIGEST)


          # check if image is attestation one, skip the clamav scan in such case

          if [[ $imageanddigest == *.att ]]

          then
              echo "$imageanddigest is an attestation image. Skipping ClamAV scan."
              exit 0
          fi

          mkdir content

          cd content

          echo Extracting image.

          if ! oc image extract --registry-config ~/.docker/config.json
          $imageanddigest; then
            echo "Unable to extract image. Skipping ClamAV scan!"
            exit 0
          fi

          echo Extraction done.

          clamscan -ri --max-scansize=250M | tee
          /tekton/home/clamscan-result.log

          echo "Executed-on: Scan was executed on version - $(clamscan
          --version)" | tee -a /tekton/home/clamscan-result.log
        securityContext:
          runAsUser: 1000
        volumeMounts:
          - mountPath: /var/lib/clamav
            name: dbfolder
          - mountPath: /work
            name: work
        workingDir: /work
      - image: >-
          quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
        name: modify-clam-output-to-json
        resources: {}
        script: >
          #!/usr/bin/env python3.9

          import json

          import dateutil.parser as parser

          import os


          clamscan_result = "/tekton/home/clamscan-result.log"

          if not os.path.exists(clamscan_result) or
          os.stat(clamscan_result).st_size == 0:
              print("clamscan-result.log file is empty, so compiled code not extracted. Parsing skipped.")
              exit(0)

          with open(clamscan_result, "r") as file:
              clam_result_str = file.read()

          def clam_result_str_to_json(clam_result_str):

              clam_result_list = clam_result_str.split("\n")
              clam_result_list.remove('')

              results_marker = \
                  clam_result_list.index("----------- SCAN SUMMARY -----------")

              hit_list = clam_result_list[:results_marker]
              summary_list = clam_result_list[(results_marker + 1):]

              r_dict = { "hits": hit_list }
              for item in summary_list:
                  # in case of blank lines
                  if not item:
                      continue
                  split_index = [c == ':' for c in item].index(True)
                  key = item[:split_index].lower()
                  key = key.replace(" ", "_")
                  value = item[(split_index + 1):].strip(" ")
                  if (key == "start_date" or key == "end_date"):
                    isodate = parser.parse(value)
                    value = isodate.isoformat()
                  r_dict[key] = value
              print(json.dumps(r_dict))
              with open('/tekton/home/clamscan-result.json', 'w') as f:
                print(json.dumps(r_dict), file=f)

          def main():
              clam_result_str_to_json(clam_result_str)

          if __name__ == "__main__":
              main()
      - image: >-
          quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
        name: store-hacbs-test-output-result
        resources: {}
        script: >
          #!/usr/bin/env bash

          source /utils.sh


          if [ -f /tekton/home/clamscan-result.json ];

          then
            cat /tekton/home/clamscan-result.json
            INFECTED_FILES=$(jq -r '.infected_files' /tekton/home/clamscan-result.json || true )
            if [ -z "${INFECTED_FILES}" ]; then
              echo "Failed to get number of infected files."
              note="Task clamav-scan failed: Unable to get number of infected files from /tekton/home/clamscan-result.json. For details, check Tekton task log."
            else
              if [[ "${INFECTED_FILES}" -gt 0 ]]; then RES="FAILURE"; else RES="SUCCESS"; fi
              note="Task clamav-scan completed: Check result for antivirus scan result."
              TEST_OUTPUT=$(make_result_json -r "${RES}" -s 1 -f "${INFECTED_FILES}" -t "$note")
            fi
          else
            note="Task clamav-scan failed: /tekton/home/clamscan-result.json doesn't exist. For details, check Tekton task log."
          fi


          ERROR_OUTPUT=$(make_result_json -r "ERROR" -t "$note")

          echo "${TEST_OUTPUT:-${ERROR_OUTPUT}}" | tee
          /tekton/results/TEST_OUTPUT
    volumes:
      - name: dbfolder
      - name: work
`

const trClairYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipelinesascode.tekton.dev/state: started
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipeline.tekton.dev/release: b8ad1b2
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    tekton.dev/tags: 'appstudio, hacbs'
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/ac8c1a7c-23de-47b7-8195-4b253cff7663
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    results.tekton.dev/log: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/logs/198203f3-9a98-30ba-bc59-299831ddd558
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    results.tekton.dev/childReadyForDeletion: 'true'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: human-resources-on-pull-request-rlrj8-clair-scan
  uid: ac8c1a7c-23de-47b7-8195-4b253cff7663
  creationTimestamp: '2023-08-25T14:58:25Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipeline.tekton.dev/release': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:tekton.dev/tags': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:tekton.dev/pipelines.minVersion': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"4ceaef28-9295-4b0b-93fc-1a03d3feadd5"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T14:58:25Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T14:59:41Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskResults': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:steps': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T14:59:41Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T15:05:55Z'
  namespace: rhtapuser-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: human-resources-on-pull-request-rlrj8
      uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  finalizers:
    - chains.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: started
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: human-resources
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    tekton.dev/task: clair-scan
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    tekton.dev/pipelineTask: clair-scan
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    tekton.dev/pipelineRun: human-resources-on-pull-request-rlrj8
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  params:
    - name: image-digest
      value: 'sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d'
    - name: image-url
      value: >-
        quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: clair-scan
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/task-clair-scan:0.1@sha256:c5602d9d6dd797da98e98fde8471ea55a788c30f74f2192807910ce5436e9b66
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
status:
  completionTime: '2023-08-25T14:59:41Z'
  conditions:
    - lastTransitionTime: '2023-08-25T14:59:41Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: human-resources-on-pull-request-rlrj8-clair-scan-pod
  startTime: '2023-08-25T14:59:25Z'
  steps:
    - container: step-get-vulnerabilities
      imageID: >-
        quay.io/redhat-appstudio/clair-in-ci@sha256:ff09557845e2ccb555fcce534e27053976260ebd11c984e3c06d2062bec336e1
      name: get-vulnerabilities
      terminated:
        containerID: >-
          cri-o://3e1116ef2e6822f8ccafdd228050ad8331098267b1465f75be97fb2d9bd94a50
        exitCode: 0
        finishedAt: '2023-08-25T14:59:40Z'
        reason: Completed
        startedAt: '2023-08-25T14:59:30Z'
    - container: step-conftest-vulnerabilities
      imageID: >-
        quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
      name: conftest-vulnerabilities
      terminated:
        containerID: >-
          cri-o://b74915dd54a8815d854d41bfd7853947e77e95cedf7de92ab5c746c142ad20e7
        exitCode: 0
        finishedAt: '2023-08-25T14:59:41Z'
        reason: Completed
        startedAt: '2023-08-25T14:59:41Z'
    - container: step-test-format-result
      imageID: >-
        quay.io/redhat-appstudio/hacbs-test@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
      name: test-format-result
      terminated:
        containerID: >-
          cri-o://e11d42d22fef1982c79a6914661f4a02d2d8bab0d24ac54d82058a01dfe2388a
        exitCode: 0
        finishedAt: '2023-08-25T14:59:41Z'
        message: >-
          [{"key":"CLAIR_SCAN_RESULT","value":"{\"vulnerabilities\":{\"critical\":0,\"high\":2,\"medium\":12,\"low\":3}}\n","type":1},{"key":"TEST_OUTPUT","value":"{\"result\":\"SUCCESS\",\"timestamp\":\"1692975581\",\"note\":\"Task
          clair-scan completed: Refer to Tekton task result CLAIR_SCAN_RESULT
          for vulnerabilities scanned by
          Clair.\",\"namespace\":\"default\",\"successes\":0,\"failures\":0,\"warnings\":0}\n","type":1}]
        reason: Completed
        startedAt: '2023-08-25T14:59:41Z'
  taskResults:
    - name: CLAIR_SCAN_RESULT
      type: string
      value: |
        {"vulnerabilities":{"critical":0,"high":2,"medium":12,"low":3}}
    - name: TEST_OUTPUT
      type: string
      value: >
        {"result":"SUCCESS","timestamp":"1692975581","note":"Task clair-scan
        completed: Refer to Tekton task result CLAIR_SCAN_RESULT for
        vulnerabilities scanned by
        Clair.","namespace":"default","successes":0,"failures":0,"warnings":0}
  taskSpec:
    description: >-
      Scans container images for vulnerabilities using Clair, by comparing the
      components of container image against Clair's vulnerability databases.
    params:
      - description: Image digest to scan.
        name: image-digest
        type: string
      - description: Image URL.
        name: image-url
        type: string
      - default: ''
        description: 'unused, should be removed in next task version.'
        name: docker-auth
        type: string
    results:
      - description: Tekton task test output.
        name: TEST_OUTPUT
        type: string
      - description: Clair scan result.
        name: CLAIR_SCAN_RESULT
        type: string
    steps:
      - env:
          - name: IMAGE_URL
            value: >-
              quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          - name: IMAGE_DIGEST
            value: >-
              sha256:dd1f97370def2173b4fcdb0b75291ed6ecdf3346a2839f6ac7f5eca3a77df37d
        image: 'quay.io/redhat-appstudio/clair-in-ci:latest'
        imagePullPolicy: Always
        name: get-vulnerabilities
        resources: {}
        script: >
          #!/usr/bin/env bash


          imagewithouttag=$(echo $IMAGE_URL | sed "s/\(.*\):.*/\1/" | tr -d
          '\n')

          # strip new-line escape symbol from parameter and save it to variable

          imageanddigest=$(echo $imagewithouttag@$IMAGE_DIGEST)


          clair-action report --image-ref=$imageanddigest
          --db-path=/tmp/matcher.db --format=quay | tee
          /tekton/home/clair-result.json || true
      - image: >-
          quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
        name: conftest-vulnerabilities
        resources: {}
        script: |
          if [ ! -s /tekton/home/clair-result.json ]; then
            echo "Previous step [get-vulnerabilities] failed: /tekton/home/clair-result.json is empty."
          else
            /usr/bin/conftest test --no-fail /tekton/home/clair-result.json \
            --policy /project/clair/vulnerabilities-check.rego --namespace required_checks \
            --output=json | tee /tekton/home/clair-vulnerabilities.json || true
          fi
        securityContext:
          capabilities:
            add:
              - SETFCAP
      - image: >-
          quay.io/redhat-appstudio/hacbs-test:v1.1.1@sha256:81acb2ba5e819b7d155ced648e48161e5f7e2bae5c0e4a0bab196651a9044afe
        name: test-format-result
        resources: {}
        script: >
          #!/usr/bin/env bash

          . /utils.sh


          if [[ ! -f /tekton/home/clair-vulnerabilities.json ]]; then
            note="Task clair-scan failed: /tekton/home/clair-vulnerabilities.json did not generate. For details, check Tekton task log."
            TEST_OUTPUT=$(make_result_json -r "ERROR" -t "$note")
            echo "/tekton/home/clair-vulnerabilities.json did not generate correctly. For details, check conftest command in Tekton task log."
            echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
            exit 0
          fi


          jq -rce \
            '{vulnerabilities:{
                critical: (.[] | .warnings? // [] | map(select(.metadata.details.name=="clair_critical_vulnerabilities").metadata."vulnerabilities_number" // 0)| add // 0),
                high: (.[] | .warnings? // [] | map(select(.metadata.details.name=="clair_high_vulnerabilities").metadata."vulnerabilities_number" // 0)| add // 0),
                medium: (.[] | .warnings? // [] | map(select(.metadata.details.name=="clair_medium_vulnerabilities").metadata."vulnerabilities_number" // 0)| add // 0),
                low: (.[] | .warnings? // [] | map(select(.metadata.details.name=="clair_low_vulnerabilities").metadata."vulnerabilities_number" // 0)| add // 0)
              }}' /tekton/home/clair-vulnerabilities.json | tee /tekton/results/CLAIR_SCAN_RESULT

          note="Task clair-scan completed: Refer to Tekton task result
          CLAIR_SCAN_RESULT for vulnerabilities scanned by Clair."

          TEST_OUTPUT=$(make_result_json -r "SUCCESS" -t "$note")

          echo "${TEST_OUTPUT}" | tee /tekton/results/TEST_OUTPUT
`

const trSummaryYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipelinesascode.tekton.dev/state: started
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipeline.tekton.dev/release: b8ad1b2
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    tekton.dev/tags: 'appstudio, hacbs'
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/b10fe769-f3e7-4f55-986f-dbe4f9b1cb11
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    results.tekton.dev/log: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/logs/85aac6d5-ab61-3bbc-82b4-ddb13dc6bc8a
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    results.tekton.dev/childReadyForDeletion: 'true'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: human-resources-on-pull-request-rlrj8-show-summary
  uid: b10fe769-f3e7-4f55-986f-dbe4f9b1cb11
  creationTimestamp: '2023-08-25T15:01:02Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipeline.tekton.dev/release': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:tekton.dev/tags': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:tekton.dev/pipelines.minVersion': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"4ceaef28-9295-4b0b-93fc-1a03d3feadd5"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T15:01:02Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T15:01:08Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:steps': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T15:01:08Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T15:06:07Z'
  namespace: rhtapuser-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: human-resources-on-pull-request-rlrj8
      uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  finalizers:
    - chains.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: started
    tekton.dev/memberOf: finally
    appstudio.openshift.io/component: human-resources
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    tekton.dev/task: summary
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    tekton.dev/pipelineTask: show-summary
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    tekton.dev/pipelineRun: human-resources-on-pull-request-rlrj8
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  params:
    - name: pipelinerun-name
      value: human-resources-on-pull-request-rlrj8
    - name: git-url
      value: >-
        https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    - name: image-url
      value: >-
        quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    - name: build-task-status
      value: Succeeded
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: summary
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/task-summary:0.1@sha256:e69f53a3991d7088d8aa2827365ab761ab7524d4269f296b4a78b0f085789d30
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
status:
  completionTime: '2023-08-25T15:01:08Z'
  conditions:
    - lastTransitionTime: '2023-08-25T15:01:08Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: human-resources-on-pull-request-rlrj8-show-summary-pod
  startTime: '2023-08-25T15:01:02Z'
  steps:
    - container: step-appstudio-summary
      imageID: >-
        registry.access.redhat.com/ubi9/ubi-minimal@sha256:dc02c6aa8199beb8ed13312d7116a94aa87b5412886bbe358845d3f0626c0f1e
      name: appstudio-summary
      terminated:
        containerID: >-
          cri-o://aa6b8e16a118ea8f712a46a380d40696d5975f7f73472f0d5276f45a20c311dc
        exitCode: 0
        finishedAt: '2023-08-25T15:01:07Z'
        reason: Completed
        startedAt: '2023-08-25T15:01:07Z'
  taskSpec:
    description: >-
      Summary Pipeline Task. Prints PipelineRun information, removes image
      repository secret used by the PipelineRun.
    params:
      - description: pipeline-run to annotate
        name: pipelinerun-name
        type: string
      - description: Git URL
        name: git-url
        type: string
      - description: Image URL
        name: image-url
        type: string
      - default: Succeeded
        description: State of build task in pipelineRun
        name: build-task-status
        type: string
    steps:
      - env:
          - name: GIT_URL
            value: >-
              https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          - name: IMAGE_URL
            value: >-
              quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
          - name: PIPELINERUN_NAME
            value: human-resources-on-pull-request-rlrj8
          - name: BUILD_TASK_STATUS
            value: Succeeded
        image: >-
          registry.access.redhat.com/ubi9/ubi-minimal:9.2-717@sha256:dc02c6aa8199beb8ed13312d7116a94aa87b5412886bbe358845d3f0626c0f1e
        name: appstudio-summary
        resources: {}
        script: |
          #!/usr/bin/env bash
          echo
          echo "Build Summary:"
          echo
          echo "Build repository: $GIT_URL"
          if [ "$BUILD_TASK_STATUS" == "Succeeded" ]; then
            echo "Generated Image is in : $IMAGE_URL"
          fi
          echo
          echo End Summary
`

const trShowSbomYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipelinesascode.tekton.dev/state: started
    pipelinesascode.tekton.dev/on-target-branch: '[main]'
    pipeline.tekton.dev/release: b8ad1b2
    pipelinesascode.tekton.dev/repo-url: 'https://github.com/jeff-phillips-18/human-resources'
    pipelinesascode.tekton.dev/sha-title: Update RHTAP references
    pipelinesascode.tekton.dev/sender: 'red-hat-trusted-app-pipeline[bot]'
    tekton.dev/tags: 'appstudio, hacbs'
    results.tekton.dev/record: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/records/ec2d790d-4489-4b89-87c9-cbc8f6d7d519
    pipelinesascode.tekton.dev/git-auth-secret: pac-gitauth-pwvj
    results.tekton.dev/log: >-
      rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5/logs/c39b89d7-b2db-3d8b-8110-1fe8cfdd3961
    build.appstudio.openshift.io/repo: >-
      https://github.com/jeff-phillips-18/human-resources?rev=b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    results.tekton.dev/result: rhtapuser-tenant/results/4ceaef28-9295-4b0b-93fc-1a03d3feadd5
    pipelinesascode.tekton.dev/log-url: 'https://console.dev.redhat.com/preview/application-pipeline'
    build.appstudio.redhat.com/target_branch: main
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    pipelinesascode.tekton.dev/max-keep-runs: '3'
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    build.appstudio.redhat.com/pull_request_number: '18'
    pipelinesascode.tekton.dev/pull-request: '18'
    results.tekton.dev/childReadyForDeletion: 'true'
    pipelinesascode.tekton.dev/url-repository: human-resources
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/sha-url: >-
      https://github.com/jeff-phillips-18/human-resources/commit/b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/on-event: '[pull_request]'
    pipelinesascode.tekton.dev/installation-id: '39596769'
    pipelinesascode.tekton.dev/event-type: pull_request
    build.appstudio.redhat.com/commit_sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  name: human-resources-on-pull-request-rlrj8-show-sbom
  uid: ec2d790d-4489-4b89-87c9-cbc8f6d7d519
  creationTimestamp: '2023-08-25T15:01:02Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:pipelinesascode.tekton.dev/on-event': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/log-url': {}
            'f:pipelinesascode.tekton.dev/max-keep-runs': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:build.appstudio.redhat.com/pull_request_number': {}
            'f:pipelinesascode.tekton.dev/on-target-branch': {}
            .: {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipeline.tekton.dev/release': {}
            'f:pipelinesascode.tekton.dev/repo-url': {}
            'f:pipelinesascode.tekton.dev/sha-url': {}
            'f:pipelinesascode.tekton.dev/installation-id': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:tekton.dev/tags': {}
            'f:pipelinesascode.tekton.dev/sha-title': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:pipelinesascode.tekton.dev/git-auth-secret': {}
            'f:tekton.dev/pipelines.minVersion': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelinesascode.tekton.dev/pull-request': {}
            'f:pipelinesascode.tekton.dev/url-repository': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:pipelinesascode.tekton.dev/repository': {}
            'f:app.kubernetes.io/managed-by': {}
            'f:appstudio.openshift.io/application': {}
            'f:pipelinesascode.tekton.dev/url-org': {}
            'f:pipelinesascode.tekton.dev/git-provider': {}
            'f:pipelinesascode.tekton.dev/event-type': {}
            'f:pipelinesascode.tekton.dev/original-prname': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:pipelinesascode.tekton.dev/sha': {}
            'f:pipelinesascode.tekton.dev/sender': {}
            'f:pipelinesascode.tekton.dev/state': {}
            'f:appstudio.openshift.io/component': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
            'f:pipelinesascode.tekton.dev/branch': {}
            'f:pipelinesascode.tekton.dev/check-run-id': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"4ceaef28-9295-4b0b-93fc-1a03d3feadd5"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-25T15:01:02Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-25T15:01:09Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:steps': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-25T15:01:09Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-25T15:06:05Z'
  namespace: rhtapuser-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: human-resources-on-pull-request-rlrj8
      uid: 4ceaef28-9295-4b0b-93fc-1a03d3feadd5
  finalizers:
    - chains.tekton.dev
  labels:
    pipelinesascode.tekton.dev/state: started
    tekton.dev/memberOf: finally
    appstudio.openshift.io/component: human-resources
    pipelinesascode.tekton.dev/sender: red-hat-trusted-app-pipeline__bot
    app.kubernetes.io/version: v0.19.4
    tekton.dev/pipeline: human-resources-on-pull-request-rlrj8
    app.kubernetes.io/managed-by: pipelinesascode.tekton.dev
    pipelinesascode.tekton.dev/check-run-id: '16217176699'
    pipelinesascode.tekton.dev/branch: main
    appstudio.openshift.io/application: test-application
    tekton.dev/task: show-sbom
    pipelinesascode.tekton.dev/url-org: jeff-phillips-18
    tekton.dev/pipelineTask: show-sbom
    pipelinesascode.tekton.dev/original-prname: human-resources-on-pull-request
    pipelinesascode.tekton.dev/pull-request: '18'
    pipelines.appstudio.openshift.io/type: build
    pipelinesascode.tekton.dev/url-repository: human-resources
    tekton.dev/pipelineRun: human-resources-on-pull-request-rlrj8
    pipelinesascode.tekton.dev/repository: human-resources
    pipelinesascode.tekton.dev/sha: b9651320a8c6fdd24cf35d2a1f8ea0a245235442
    pipelinesascode.tekton.dev/git-provider: github
    pipelinesascode.tekton.dev/event-type: pull_request
spec:
  params:
    - name: IMAGE_URL
      value: >-
        quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: show-sbom
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/task-show-sbom:0.1@sha256:7db0af43dcebaeb33e34413148370e17078c30fd2fc78fb84c8941b444199f36
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
status:
  completionTime: '2023-08-25T15:01:09Z'
  conditions:
    - lastTransitionTime: '2023-08-25T15:01:09Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: human-resources-on-pull-request-rlrj8-show-sbom-pod
  startTime: '2023-08-25T15:01:02Z'
  steps:
    - container: step-show-sbom
      imageID: >-
        quay.io/redhat-appstudio/cosign@sha256:c883d6f8d39148f2cea71bff4622d196d89df3e510f36c140c097b932f0dd5d5
      name: show-sbom
      terminated:
        containerID: >-
          cri-o://6c057c683644abb99da528594de18b2f7d56244c44d04e1c92f6eba1e0e083e7
        exitCode: 0
        finishedAt: '2023-08-25T15:01:08Z'
        reason: Completed
        startedAt: '2023-08-25T15:01:07Z'
  taskSpec:
    description: >-
      Shows the Software Bill of Materials (SBOM) generated for the built image
      in CyloneDX JSON format.
    params:
      - description: Fully qualified image name to show SBOM for.
        name: IMAGE_URL
        type: string
    steps:
      - env:
          - name: IMAGE_URL
            value: >-
              quay.io/redhat-user-workloads-stage/rhtapuser-tenant/test-application/human-resources:on-pr-b9651320a8c6fdd24cf35d2a1f8ea0a245235442
        image: 'quay.io/redhat-appstudio/cosign:v2.1.1'
        name: show-sbom
        resources: {}
        script: |
          #!/busybox/sh
          cosign download sbom $IMAGE_URL 2>err
          RET=$?
          if [ $RET -ne 0 ]; then
            echo Failed to get SBOM >&2
            cat err >&2
          fi
          exit $RET
`

const tooBigNumPRYaml = `
apiVersion: tekton.dev/v1beta1
kind: PipelineRun
metadata:
  generateName: devfile-sample-
  annotations:
    appstudio.openshift.io/snapshot: load-app-7089021e-3585-438c-85e8-0a1ec8381a9f-9thsf
    results.tekton.dev/record: >-
      test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45/records/579cdd9f-5dcc-436c-ab9b-eb770530ff45
    build.appstudio.redhat.com/pipeline_name: docker-build
    build.appstudio.openshift.io/repo: >-
      https://github.com/nodeshift-starters/devfile-sample?rev=cb64992eebf5c18900d283a9bca08b4ab5db2874
    chains.tekton.dev/signed: 'true'
    results.tekton.dev/result: test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45
    build.appstudio.redhat.com/target_branch: main
    results.tekton.dev/childReadyForDeletion: 'true'
    build.appstudio.redhat.com/bundle: >-
      quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:docker-build-124dd76239d17a3bdb936ed26e403c9c9c4e947b
    build.appstudio.redhat.com/commit_sha: cb64992eebf5c18900d283a9bca08b4ab5db2874
  name: devfile-sample-nsj6m
  uid: 579cdd9f-5dcc-436c-ab9b-eb770530ff45
  creationTimestamp: '2023-08-30T19:00:02Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:labels':
            'f:pipelines.openshift.io/runtime': {}
            'f:pipelines.openshift.io/strategy': {}
            'f:pipelines.openshift.io/used-by': {}
            'f:tekton.dev/pipeline': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-30T19:00:02Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          'f:finallyStartTime': {}
          'f:pipelineResults': {}
          'f:conditions': {}
          .: {}
          'f:childReferences': {}
          'f:pipelineSpec':
            .: {}
            'f:finally': {}
            'f:params': {}
            'f:results': {}
            'f:tasks': {}
            'f:workspaces': {}
          'f:skippedTasks': {}
          'f:startTime': {}
          'f:completionTime': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-30T19:01:36Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev/pipelinerun"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-30T19:01:37Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            .: {}
            'f:appstudio.openshift.io/snapshot': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:build.appstudio.redhat.com/bundle': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            'f:build.appstudio.redhat.com/pipeline_name': {}
            'f:build.appstudio.redhat.com/target_branch': {}
          'f:generateName': {}
          'f:labels':
            .: {}
            'f:appstudio.openshift.io/application': {}
            'f:appstudio.openshift.io/component': {}
            'f:pipelines.appstudio.openshift.io/type': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"efeffdc5-984c-43ec-9707-65ebccc85011"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:pipelineRef':
            .: {}
            'f:bundle': {}
            'f:name': {}
          'f:workspaces': {}
      manager: manager
      operation: Update
      time: '2023-08-30T19:01:38Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/record': {}
            'f:results.tekton.dev/result': {}
      manager: watcher
      operation: Update
      time: '2023-08-30T19:01:41Z'
  namespace: test-rhtap-95-tenant
  ownerReferences:
    - apiVersion: appstudio.redhat.com/v1alpha1
      kind: Component
      name: devfile-sample
      uid: efeffdc5-984c-43ec-9707-65ebccc85011
  finalizers:
    - chains.tekton.dev/pipelinerun
  labels:
    appstudio.openshift.io/application: load-app-7089021e-3585-438c-85e8-0a1ec8381a9f
    appstudio.openshift.io/component: devfile-sample
    pipelines.appstudio.openshift.io/type: build
    pipelines.openshift.io/runtime: generic
    pipelines.openshift.io/strategy: docker
    pipelines.openshift.io/used-by: build-cloud
    tekton.dev/pipeline: docker-build
spec:
  params:
    - name: dockerfile
      value: Dockerfile
    - name: git-url
      value: 'https://github.com/nodeshift-starters/devfile-sample.git'
    - name: output-image
      value: >-
        quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
    - name: path-context
      value: .
    - name: revision
      value: main
    - name: skip-checks
      value: 'true'
  pipelineRef:
    bundle: >-
      quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:docker-build-124dd76239d17a3bdb936ed26e403c9c9c4e947b
    name: docker-build
  serviceAccountName: appstudio-pipeline
  timeout: 1h0m0s
  workspaces:
    - name: workspace
      volumeClaimTemplate:
        metadata:
          creationTimestamp: null
        spec:
          accessModes:
            - ReadWriteOnce
          resources:
            requests:
              storage: 1Gi
        status: {}
status:
  childReferences:
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: devfile-sample-nsj6m-init
      pipelineTaskName: init
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: devfile-sample-nsj6m-clone-repository
      pipelineTaskName: clone-repository
      whenExpressions:
        - input: $(tasks.init.results.build)
          operator: in
          values:
            - 'true'
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: devfile-sample-nsj6m-build-container
      pipelineTaskName: build-container
      whenExpressions:
        - input: $(tasks.init.results.build)
          operator: in
          values:
            - 'true'
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: devfile-sample-nsj6m-show-sbom
      pipelineTaskName: show-sbom
    - apiVersion: tekton.dev/v1beta1
      kind: TaskRun
      name: devfile-sample-nsj6m-show-summary
      pipelineTaskName: show-summary
  completionTime: '2023-08-30T19:01:36Z'
  conditions:
    - lastTransitionTime: '2023-08-30T19:01:36Z'
      message: 'Tasks Completed: 5 (Failed: 0, Cancelled 0), Skipped: 7'
      reason: Completed
      status: 'True'
      type: Succeeded
  finallyStartTime: '2023-08-30T19:01:30Z'
  pipelineResults:
    - name: IMAGE_URL
      value: >-
        quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
    - name: IMAGE_DIGEST
      value: 'sha256:9cc2777884b232b63f67c352ba96866edac8cf35fc14a87107bfcf64faaff7db'
    - name: CHAINS-GIT_URL
      value: 'https://github.com/nodeshift-starters/devfile-sample.git'
    - name: CHAINS-GIT_COMMIT
      value: cb64992eebf5c18900d283a9bca08b4ab5db2874
    - name: JAVA_COMMUNITY_DEPENDENCIES
      value: ''
  pipelineSpec:
    finally:
      - name: show-sbom
        params:
          - name: IMAGE_URL
            value: $(tasks.build-container.results.IMAGE_URL)
        taskRef:
          kind: Task
          params:
            - name: name
              value: show-sbom
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:show-sbom-0.1@sha256:7db0af43dcebaeb33e34413148370e17078c30fd2fc78fb84c8941b444199f36
            - name: kind
              value: task
          resolver: bundles
        status:
          completionTime: '2023-08-30T19:01:36Z'
          conditions:
            - lastTransitionTime: '2023-08-30T19:01:36Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: devfile-sample-nsj6m-show-sbom-pod
          startTime: '2023-08-30T19:01:30Z'
          steps:
            - container: step-show-sbom
              imageID: >-
                quay.io/redhat-appstudio/cosign@sha256:c883d6f8d39148f2cea71bff4622d196d89df3e510f36c140c097b932f0dd5d5
              name: show-sbom
              terminated:
                containerID: >-
                  cri-o://6b18d743bb70bdc2ccb2dbb37d80571ed897fca535ace395329032e68983abd7
                exitCode: 0
                finishedAt: '2023-08-30T19:01:36Z'
                reason: Completed
                startedAt: '2023-08-30T19:01:35Z'
          taskSpec:
            description: >-
              Shows the Software Bill of Materials (SBOM) generated for the
              built image in CyloneDX JSON format.
            params:
              - description: Fully qualified image name to show SBOM for.
                name: IMAGE_URL
                type: string
            steps:
              - env:
                  - name: IMAGE_URL
                    value: >-
                      quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
                image: 'quay.io/redhat-appstudio/cosign:v2.1.1'
                name: show-sbom
                resources: {}
                script: |
                  #!/busybox/sh
                  cosign download sbom $IMAGE_URL 2>err
                  RET=$?
                  if [ $RET -ne 0 ]; then
                    echo Failed to get SBOM >&2
                    cat err >&2
                  fi
                  exit $RET
          duration: 6s
          reason: Succeeded
      - name: show-summary
        params:
          - name: pipelinerun-name
            value: devfile-sample-nsj6m
          - name: git-url
            value: >-
              $(tasks.clone-repository.results.url)?rev=$(tasks.clone-repository.results.commit)
          - name: image-url
            value: >-
              quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
          - name: build-task-status
            value: $(tasks.build-container.status)
        taskRef:
          kind: Task
          params:
            - name: name
              value: summary
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:summary-0.1@sha256:e69f53a3991d7088d8aa2827365ab761ab7524d4269f296b4a78b0f085789d30
            - name: kind
              value: task
          resolver: bundles
        status:
          completionTime: '2023-08-30T19:01:36Z'
          conditions:
            - lastTransitionTime: '2023-08-30T19:01:36Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: devfile-sample-nsj6m-show-summary-pod
          startTime: '2023-08-30T19:01:30Z'
          steps:
            - container: step-appstudio-summary
              imageID: >-
                registry.access.redhat.com/ubi9/ubi-minimal@sha256:dc02c6aa8199beb8ed13312d7116a94aa87b5412886bbe358845d3f0626c0f1e
              name: appstudio-summary
              terminated:
                containerID: >-
                  cri-o://149f69bd03f32d573735805c98eb57f5ef7de8dd4e91a207a084bbe1f60b907d
                exitCode: 0
                finishedAt: '2023-08-30T19:01:35Z'
                reason: Completed
                startedAt: '2023-08-30T19:01:35Z'
          taskSpec:
            description: >-
              Summary Pipeline Task. Prints PipelineRun information, removes
              image repository secret used by the PipelineRun.
            params:
              - description: pipeline-run to annotate
                name: pipelinerun-name
                type: string
              - description: Git URL
                name: git-url
                type: string
              - description: Image URL
                name: image-url
                type: string
              - default: Succeeded
                description: State of build task in pipelineRun
                name: build-task-status
                type: string
            steps:
              - env:
                  - name: GIT_URL
                    value: >-
                      https://github.com/nodeshift-starters/devfile-sample.git?rev=cb64992eebf5c18900d283a9bca08b4ab5db2874
                  - name: IMAGE_URL
                    value: >-
                      quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
                  - name: PIPELINERUN_NAME
                    value: devfile-sample-nsj6m
                  - name: BUILD_TASK_STATUS
                    value: Succeeded
                image: >-
                  registry.access.redhat.com/ubi9/ubi-minimal:9.2-717@sha256:dc02c6aa8199beb8ed13312d7116a94aa87b5412886bbe358845d3f0626c0f1e
                name: appstudio-summary
                resources: {}
                script: |
                  #!/usr/bin/env bash
                  echo
                  echo "Build Summary:"
                  echo
                  echo "Build repository: $GIT_URL"
                  if [ "$BUILD_TASK_STATUS" == "Succeeded" ]; then
                    echo "Generated Image is in : $IMAGE_URL"
                  fi
                  echo
                  echo End Summary
          duration: 6s
          reason: Succeeded
    params:
      - description: Source Repository URL
        name: git-url
        type: string
      - default: ''
        description: Revision of the Source Repository
        name: revision
        type: string
      - description: Fully Qualified Output Image
        name: output-image
        type: string
      - default: .
        description: The path to your source code
        name: path-context
        type: string
      - default: Dockerfile
        description: Path to the Dockerfile
        name: dockerfile
        type: string
      - default: 'false'
        description: Force rebuild image
        name: rebuild
        type: string
      - default: 'false'
        description: Skip checks against built image
        name: skip-checks
        type: string
      - default: 'true'
        description: 'Skip optional checks, set false if you want to run optional checks'
        name: skip-optional
        type: string
      - default: 'false'
        description: Execute the build with network isolation
        name: hermetic
        type: string
      - default: ''
        description: Build dependencies to be prefetched by Cachi2
        name: prefetch-input
        type: string
      - default: 'false'
        description: Java build
        name: java
        type: string
      - default: ''
        description: >-
          Image tag expiration time, time values could be something like 1h, 2d,
          3w for hours, days, and weeks, respectively.
        name: image-expires-after
        type: string
    results:
      - description: ''
        name: IMAGE_URL
        value: $(tasks.build-container.results.IMAGE_URL)
      - description: ''
        name: IMAGE_DIGEST
        value: $(tasks.build-container.results.IMAGE_DIGEST)
      - description: ''
        name: CHAINS-GIT_URL
        value: $(tasks.clone-repository.results.url)
      - description: ''
        name: CHAINS-GIT_COMMIT
        value: $(tasks.clone-repository.results.commit)
      - description: ''
        name: JAVA_COMMUNITY_DEPENDENCIES
        value: $(tasks.build-container.results.JAVA_COMMUNITY_DEPENDENCIES)
    tasks:
      - name: init
        params:
          - name: image-url
            value: >-
              quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
          - name: rebuild
            value: 'false'
          - name: skip-checks
            value: 'true'
          - name: skip-optional
            value: 'true'
          - name: pipelinerun-name
            value: devfile-sample-nsj6m
          - name: pipelinerun-uid
            value: 579cdd9f-5dcc-436c-ab9b-eb770530ff45
        taskRef:
          kind: Task
          params:
            - name: name
              value: init
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:init-0.1@sha256:26586a7ef08c3e86dfdaf0a5cc38dd3d70c4c02db1331b469caaed0a0f5b3d86
            - name: kind
              value: task
          resolver: bundles
        status:
          completionTime: '2023-08-30T19:00:09Z'
          conditions:
            - lastTransitionTime: '2023-08-30T19:00:09Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: devfile-sample-nsj6m-init-pod
          startTime: '2023-08-30T19:00:04Z'
          steps:
            - container: step-init
              imageID: >-
                registry.redhat.io/openshift4/ose-cli@sha256:73df37794ffff7de1101016c23dc623e4990810390ebdabcbbfa065214352c7c
              name: init
              terminated:
                containerID: >-
                  cri-o://916f5f33bbb4819b281ba24ab276c20524c3ff33ad0fa86c213924255cadf59f
                exitCode: 0
                finishedAt: '2023-08-30T19:00:08Z'
                message: >-
                  [{"key":"build","value":"true","type":1},{"key":"container-registry-secret","value":"unused\n","type":1}]
                reason: Completed
                startedAt: '2023-08-30T19:00:08Z'
          taskResults:
            - name: build
              type: string
              value: 'true'
            - name: container-registry-secret
              type: string
              value: |
                unused
          taskSpec:
            description: >-
              Initialize Pipeline Task, include flags for rebuild and auth.
              Generates image repository secret used by the PipelineRun.
            params:
              - description: Image URL for build by PipelineRun
                name: image-url
                type: string
              - default: 'false'
                description: Rebuild the image if exists
                name: rebuild
                type: string
              - default: 'false'
                description: Skip checks against built image
                name: skip-checks
                type: string
              - default: 'true'
                description: >-
                  Skip optional checks, set false if you want to run optional
                  checks
                name: skip-optional
                type: string
              - default: ''
                description: 'unused, should be removed in next task version'
                name: pipelinerun-name
                type: string
              - default: ''
                description: 'unused, should be removed in next task version'
                name: pipelinerun-uid
                type: string
            results:
              - description: Defines if the image in param image-url should be built
                name: build
                type: string
              - description: 'unused, should be removed in next task version'
                name: container-registry-secret
                type: string
            steps:
              - env:
                  - name: IMAGE_URL
                    value: >-
                      quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
                  - name: REBUILD
                    value: 'false'
                  - name: SKIP_CHECKS
                    value: 'true'
                  - name: SKIP_OPTIONAL
                    value: 'true'
                image: >-
                  registry.redhat.io/openshift4/ose-cli:4.13@sha256:73df37794ffff7de1101016c23dc623e4990810390ebdabcbbfa065214352c7c
                name: init
                resources: {}
                script: >
                  #!/bin/bash

                  echo "Build Initialize: $IMAGE_URL"

                  echo


                  echo "Determine if Image Already Exists"

                  # Build the image when image does not exists or rebuild is set
                  to true

                  if ! oc image info $IMAGE_URL &>/dev/null || [ "$REBUILD" ==
                  "true" ] || [ "$SKIP_CHECKS" == "false" ]; then
                    echo -n "true" > /tekton/results/build
                  else
                    echo -n "false" > /tekton/results/build
                  fi

                  echo unused > /tekton/results/container-registry-secret
          duration: 5s
          reason: Succeeded
      - name: clone-repository
        params:
          - name: url
            value: 'https://github.com/nodeshift-starters/devfile-sample.git'
          - name: revision
            value: main
        runAfter:
          - init
        taskRef:
          kind: Task
          params:
            - name: name
              value: git-clone
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:git-clone-0.1@sha256:1f84973a21aabea38434b1f663abc4cb2d86565a9c7aae1f90decb43a8fa48eb
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(tasks.init.results.build)
            operator: in
            values:
              - 'true'
        workspaces:
          - name: output
            workspace: workspace
          - name: basic-auth
            workspace: git-auth
        status:
          completionTime: '2023-08-30T19:00:29Z'
          conditions:
            - lastTransitionTime: '2023-08-30T19:00:29Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: devfile-sample-nsj6m-clone-repository-pod
          startTime: '2023-08-30T19:00:09Z'
          steps:
            - container: step-clone
              imageID: >-
                registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8@sha256:2fa0b06d52b04f377c696412e19307a9eff27383f81d87aae0b4f71672a1cd0b
              name: clone
              terminated:
                containerID: >-
                  cri-o://2f18fd7be23d4e6ac18bd3066b7f4efbedb0d4527f29c321ccf50e5b3187ae8d
                exitCode: 0
                finishedAt: '2023-08-30T19:00:28Z'
                message: >-
                  [{"key":"commit","value":"cb64992eebf5c18900d283a9bca08b4ab5db2874","type":1},{"key":"url","value":"https://github.com/nodeshift-starters/devfile-sample.git","type":1}]
                reason: Completed
                startedAt: '2023-08-30T19:00:27Z'
            - container: step-symlink-check
              imageID: >-
                registry.redhat.io/ubi9@sha256:089bd3b82a78ac45c0eed231bb58bfb43bfcd0560d9bba240fc6355502c92976
              name: symlink-check
              terminated:
                containerID: >-
                  cri-o://1492a8338e20bcc7897d81af16d33e47ca26a51b45e1c238ff657610b48a243a
                exitCode: 0
                finishedAt: '2023-08-30T19:00:28Z'
                message: >-
                  [{"key":"commit","value":"cb64992eebf5c18900d283a9bca08b4ab5db2874","type":1},{"key":"url","value":"https://github.com/nodeshift-starters/devfile-sample.git","type":1}]
                reason: Completed
                startedAt: '2023-08-30T19:00:28Z'
          taskResults:
            - name: commit
              type: string
              value: cb64992eebf5c18900d283a9bca08b4ab5db2874
            - name: url
              type: string
              value: 'https://github.com/nodeshift-starters/devfile-sample.git'
          taskSpec:
            description: >-
              The git-clone Task will clone a repo from the provided url into
              the output Workspace. By default the repo will be cloned into the
              root of your Workspace.
            params:
              - description: Repository URL to clone from.
                name: url
                type: string
              - default: ''
                description: 'Revision to checkout. (branch, tag, sha, ref, etc...)'
                name: revision
                type: string
              - default: ''
                description: Refspec to fetch before checking out revision.
                name: refspec
                type: string
              - default: 'true'
                description: Initialize and fetch git submodules.
                name: submodules
                type: string
              - default: '1'
                description: >-
                  Perform a shallow clone, fetching only the most recent N
                  commits.
                name: depth
                type: string
              - default: 'true'
                description: >-
                  Set the http.sslVerify global git config. Setting this to
                  false is not advised unless you are sure that you trust your
                  git remote.
                name: sslVerify
                type: string
              - default: ''
                description: >-
                  Subdirectory inside the output Workspace to clone the repo
                  into.
                name: subdirectory
                type: string
              - default: ''
                description: >-
                  Define the directory patterns to match or exclude when
                  performing a sparse checkout.
                name: sparseCheckoutDirectories
                type: string
              - default: 'true'
                description: >-
                  Clean out the contents of the destination directory if it
                  already exists before cloning.
                name: deleteExisting
                type: string
              - default: ''
                description: HTTP proxy server for non-SSL requests.
                name: httpProxy
                type: string
              - default: ''
                description: HTTPS proxy server for SSL requests.
                name: httpsProxy
                type: string
              - default: ''
                description: Opt out of proxying HTTP/HTTPS requests.
                name: noProxy
                type: string
              - default: 'true'
                description: >-
                  Log the commands that are executed during git-clone's
                  operation.
                name: verbose
                type: string
              - default: >-
                  registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8:v1.8.2-8@sha256:a538c423e7a11aae6ae582a411fdb090936458075f99af4ce5add038bb6983e8
                description: The image providing the git-init binary that this Task runs.
                name: gitInitImage
                type: string
              - default: /tekton/home
                description: >
                  Absolute path to the user's home directory. Set this
                  explicitly if you are running the image as a non-root user or
                  have overridden

                  the gitInitImage param with an image containing custom user
                  configuration.
                name: userHome
                type: string
              - default: 'true'
                description: >
                  Check symlinks in the repo. If they're pointing outside of the
                  repo, the build will fail.
                name: enableSymlinkCheck
                type: string
            results:
              - description: The precise commit SHA that was fetched by this Task.
                name: commit
                type: string
              - description: The precise URL that was fetched by this Task.
                name: url
                type: string
            steps:
              - env:
                  - name: HOME
                    value: /tekton/home
                  - name: PARAM_URL
                    value: 'https://github.com/nodeshift-starters/devfile-sample.git'
                  - name: PARAM_REVISION
                    value: main
                  - name: PARAM_REFSPEC
                  - name: PARAM_SUBMODULES
                    value: 'true'
                  - name: PARAM_DEPTH
                    value: '1'
                  - name: PARAM_SSL_VERIFY
                    value: 'true'
                  - name: PARAM_SUBDIRECTORY
                  - name: PARAM_DELETE_EXISTING
                    value: 'true'
                  - name: PARAM_HTTP_PROXY
                  - name: PARAM_HTTPS_PROXY
                  - name: PARAM_NO_PROXY
                  - name: PARAM_VERBOSE
                    value: 'true'
                  - name: PARAM_SPARSE_CHECKOUT_DIRECTORIES
                  - name: PARAM_USER_HOME
                    value: /tekton/home
                  - name: WORKSPACE_OUTPUT_PATH
                    value: /workspace/output
                  - name: WORKSPACE_SSH_DIRECTORY_BOUND
                    value: 'false'
                  - name: WORKSPACE_SSH_DIRECTORY_PATH
                  - name: WORKSPACE_BASIC_AUTH_DIRECTORY_BOUND
                    value: 'false'
                  - name: WORKSPACE_BASIC_AUTH_DIRECTORY_PATH
                image: >-
                  registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8:v1.8.2-8@sha256:a538c423e7a11aae6ae582a411fdb090936458075f99af4ce5add038bb6983e8
                name: clone
                resources: {}
                script: >
                  #!/usr/bin/env sh

                  set -eu


                  if [ "${PARAM_VERBOSE}" = "true" ] ; then
                    set -x
                  fi


                  if [ "${WORKSPACE_BASIC_AUTH_DIRECTORY_BOUND}" = "true" ] ;
                  then
                    if [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.git-credentials" ] && [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.gitconfig" ]; then
                      cp "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.git-credentials" "${PARAM_USER_HOME}/.git-credentials"
                      cp "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.gitconfig" "${PARAM_USER_HOME}/.gitconfig"
                    # Compatibility with kubernetes.io/basic-auth secrets
                    elif [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/username" ] && [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/password" ]; then
                      HOSTNAME=$(echo $PARAM_URL | awk -F/ '{print $3}')
                      echo "https://$(cat ${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/username):$(cat ${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/password)@$HOSTNAME" > "${PARAM_USER_HOME}/.git-credentials"
                      echo -e "[credential \"https://$HOSTNAME\"]\n  helper = store" > "${PARAM_USER_HOME}/.gitconfig"
                    else
                      echo "Unknown basic-auth workspace format"
                      exit 1
                    fi
                    chmod 400 "${PARAM_USER_HOME}/.git-credentials"
                    chmod 400 "${PARAM_USER_HOME}/.gitconfig"
                  fi


                  if [ "${WORKSPACE_SSH_DIRECTORY_BOUND}" = "true" ] ; then
                    cp -R "${WORKSPACE_SSH_DIRECTORY_PATH}" "${PARAM_USER_HOME}"/.ssh
                    chmod 700 "${PARAM_USER_HOME}"/.ssh
                    chmod -R 400 "${PARAM_USER_HOME}"/.ssh/*
                  fi


                  CHECKOUT_DIR="${WORKSPACE_OUTPUT_PATH}/${PARAM_SUBDIRECTORY}"


                  cleandir() {
                    # Delete any existing contents of the repo directory if it exists.
                    #
                    # We don't just "rm -rf ${CHECKOUT_DIR}" because ${CHECKOUT_DIR} might be "/"
                    # or the root of a mounted volume.
                    if [ -d "${CHECKOUT_DIR}" ] ; then
                      # Delete non-hidden files and directories
                      rm -rf "${CHECKOUT_DIR:?}"/*
                      # Delete files and directories starting with . but excluding ..
                      rm -rf "${CHECKOUT_DIR}"/.[!.]*
                      # Delete files and directories starting with .. plus any other character
                      rm -rf "${CHECKOUT_DIR}"/..?*
                    fi
                  }


                  if [ "${PARAM_DELETE_EXISTING}" = "true" ] ; then
                    cleandir
                  fi


                  test -z "${PARAM_HTTP_PROXY}" || export
                  HTTP_PROXY="${PARAM_HTTP_PROXY}"

                  test -z "${PARAM_HTTPS_PROXY}" || export
                  HTTPS_PROXY="${PARAM_HTTPS_PROXY}"

                  test -z "${PARAM_NO_PROXY}" || export
                  NO_PROXY="${PARAM_NO_PROXY}"


                  /ko-app/git-init \
                    -url="${PARAM_URL}" \
                    -revision="${PARAM_REVISION}" \
                    -refspec="${PARAM_REFSPEC}" \
                    -path="${CHECKOUT_DIR}" \
                    -sslVerify="${PARAM_SSL_VERIFY}" \
                    -submodules="${PARAM_SUBMODULES}" \
                    -depth="${PARAM_DEPTH}" \
                    -sparseCheckoutDirectories="${PARAM_SPARSE_CHECKOUT_DIRECTORIES}"
                  cd "${CHECKOUT_DIR}"

                  RESULT_SHA="$(git rev-parse HEAD)"

                  EXIT_CODE="$?"

                  if [ "${EXIT_CODE}" != 0 ] ; then
                    exit "${EXIT_CODE}"
                  fi

                  printf "%s" "${RESULT_SHA}" > "/tekton/results/commit"

                  printf "%s" "${PARAM_URL}" > "/tekton/results/url"
                securityContext:
                  runAsUser: 0
              - env:
                  - name: PARAM_ENABLE_SYMLINK_CHECK
                    value: 'true'
                  - name: PARAM_SUBDIRECTORY
                  - name: WORKSPACE_OUTPUT_PATH
                    value: /workspace/output
                image: 'registry.redhat.io/ubi9:9.2-696'
                name: symlink-check
                resources: {}
                script: |
                  #!/usr/bin/env bash
                  set -euo pipefail

                  CHECKOUT_DIR="${WORKSPACE_OUTPUT_PATH}/${PARAM_SUBDIRECTORY}"
                  check_symlinks() {
                    FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO=false
                    while read symlink
                    do
                      target=$(readlink -f "$symlink")
                      if ! [[ "$target" =~ ^$CHECKOUT_DIR ]]; then
                        echo "The cloned repository contains symlink pointing outside of the cloned repository: $symlink"
                        FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO=true
                      fi
                    done < <(find $CHECKOUT_DIR -type l -print)
                    if [ "$FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO" = true ] ; then
                      return 1
                    fi
                  }

                  if [ "${PARAM_ENABLE_SYMLINK_CHECK}" = "true" ] ; then
                    echo "Running symlink check"
                    check_symlinks
                  fi
            workspaces:
              - description: >-
                  The git repo will be cloned onto the volume backing this
                  Workspace.
                name: output
              - description: >
                  A .ssh directory with private key, known_hosts, config, etc.
                  Copied to

                  the user's home before git commands are executed. Used to
                  authenticate

                  with the git remote when performing the clone. Binding a
                  Secret to this

                  Workspace is strongly recommended over other volume types.
                name: ssh-directory
                optional: true
              - description: >
                  A Workspace containing a .gitconfig and .git-credentials file
                  or username and password.

                  These will be copied to the user's home before any git
                  commands are run. Any

                  other files in this Workspace are ignored. It is strongly
                  recommended

                  to use ssh-directory over basic-auth whenever possible and to
                  bind a

                  Secret to this Workspace over other volume types.
                name: basic-auth
                optional: true
          duration: 20s
          reason: Succeeded
      - name: prefetch-dependencies
        params:
          - name: input
            value: ''
        runAfter:
          - clone-repository
        taskRef:
          kind: Task
          params:
            - name: name
              value: prefetch-dependencies
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:prefetch-dependencies-0.1@sha256:c7b7f13d5d2a1545e95c2d56521327001d56ba54645900db41aa414607eff1e5
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'false'
            operator: in
            values:
              - 'true'
        workspaces:
          - name: source
            workspace: workspace
        status:
          reason: Skipped
      - name: build-container
        params:
          - name: IMAGE
            value: >-
              quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
          - name: DOCKERFILE
            value: Dockerfile
          - name: CONTEXT
            value: .
          - name: HERMETIC
            value: 'false'
          - name: PREFETCH_INPUT
            value: ''
          - name: IMAGE_EXPIRES_AFTER
            value: ''
          - name: COMMIT_SHA
            value: $(tasks.clone-repository.results.commit)
        runAfter:
          - prefetch-dependencies
        taskRef:
          kind: Task
          params:
            - name: name
              value: buildah
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:buildah-0.1@sha256:bfb5edabab8128e24608df608448b9392fd0a2b61ac05a53e83aa60d8929b73f
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: $(tasks.init.results.build)
            operator: in
            values:
              - 'true'
        workspaces:
          - name: source
            workspace: workspace
        status:
          completionTime: '2023-08-30T19:01:30Z'
          conditions:
            - lastTransitionTime: '2023-08-30T19:01:30Z'
              message: All Steps have completed executing
              reason: Succeeded
              status: 'True'
              type: Succeeded
          podName: devfile-sample-nsj6m-build-container-pod
          startTime: '2023-08-30T19:00:29Z'
          steps:
            - container: step-build
              imageID: >-
                quay.io/redhat-appstudio/buildah@sha256:381e9bfedd59701477621da93892106873a6951b196105d3d2d85c3f6d7b569b
              name: build
              terminated:
                containerID: >-
                  cri-o://24d9f9c59f1c0fe114d05b1c29ac7209c3d3482716cc740f4783c8903c4721fb
                exitCode: 0
                finishedAt: '2023-08-30T19:01:16Z'
                reason: Completed
                startedAt: '2023-08-30T19:00:50Z'
            - container: step-sbom-syft-generate
              imageID: >-
                quay.io/redhat-appstudio/syft@sha256:244a17ce220a0b7a54c862c4fe3b72ce92799910c5eff8e94ac2f121fa5b4a53
              name: sbom-syft-generate
              terminated:
                containerID: >-
                  cri-o://cb51f07d379e15f8a8755bfb84428ff60de6989f0a0c2a1de3fa07927b09ec2c
                exitCode: 0
                finishedAt: '2023-08-30T19:01:18Z'
                reason: Completed
                startedAt: '2023-08-30T19:01:16Z'
            - container: step-analyse-dependencies-java-sbom
              imageID: >-
                quay.io/redhat-appstudio/hacbs-jvm-build-request-processor@sha256:b198cf4b33dab59ce8ac25afd4e1001390db29ca2dec83dc8a1e21b0359ce743
              name: analyse-dependencies-java-sbom
              terminated:
                containerID: >-
                  cri-o://b7448e62a57890749446e24a4f7f51ea2fa6333d9616c1193f4f92d9728e8fe6
                exitCode: 0
                finishedAt: '2023-08-30T19:01:18Z'
                message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
                reason: Completed
                startedAt: '2023-08-30T19:01:18Z'
            - container: step-merge-syft-sboms
              imageID: >-
                registry.access.redhat.com/ubi9/python-39@sha256:562b4f9cd4e5abeba54c56fd2290096a71f6e10aa8c1c18f43c9d6962c2d4d41
              name: merge-syft-sboms
              terminated:
                containerID: >-
                  cri-o://7082cb7ddc3c629dca86aac4b5ae553eeb09960402ff503b49753b4f8c1ef7a1
                exitCode: 0
                finishedAt: '2023-08-30T19:01:19Z'
                message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
                reason: Completed
                startedAt: '2023-08-30T19:01:18Z'
            - container: step-merge-cachi2-sbom
              imageID: >-
                quay.io/redhat-appstudio/cachi2@sha256:46097f22b57e4d48a3fce96d931e08ccfe3a3e6421362d5f9353961279078eef
              name: merge-cachi2-sbom
              terminated:
                containerID: >-
                  cri-o://543053c4be3079d0d266d494b520852f15d3329a5b2e5ab36f12ffb138eac2e5
                exitCode: 0
                finishedAt: '2023-08-30T19:01:19Z'
                message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
                reason: Completed
                startedAt: '2023-08-30T19:01:19Z'
            - container: step-create-purl-sbom
              imageID: >-
                registry.access.redhat.com/ubi9/python-39@sha256:562b4f9cd4e5abeba54c56fd2290096a71f6e10aa8c1c18f43c9d6962c2d4d41
              name: create-purl-sbom
              terminated:
                containerID: >-
                  cri-o://7e41fead679f12a2291dc1c1d59c4fff229cfd7f17cebea7381aa9dd2c6472a5
                exitCode: 0
                finishedAt: '2023-08-30T19:01:19Z'
                message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
                reason: Completed
                startedAt: '2023-08-30T19:01:19Z'
            - container: step-inject-sbom-and-push
              imageID: >-
                registry.access.redhat.com/ubi9/buildah@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
              name: inject-sbom-and-push
              terminated:
                containerID: >-
                  cri-o://f17b5f541734bb1c60ba25aa733925f009345ba2b47d77ff25b5c8d28f08802b
                exitCode: 0
                finishedAt: '2023-08-30T19:01:27Z'
                message: >-
                  [{"key":"BASE_IMAGES_DIGESTS","value":"registry.access.redhat.com/ubi8/nodejs-18:latest@sha256:8a634a49c4e8d1753cddd3c047a4324d405fd2f6723f67b7ac274bb5fe72ff44\nregistry.access.redhat.com/ubi8/nodejs-18-minimal:latest@sha256:a1372e99afe18a64b9dbb70612452efe02e934db136fba00ba1da417ede408a0\n","type":1},{"key":"IMAGE_DIGEST","value":"sha256:9cc2777884b232b63f67c352ba96866edac8cf35fc14a87107bfcf64faaff7db","type":1},{"key":"IMAGE_URL","value":"quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002","type":1},{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]
                reason: Completed
                startedAt: '2023-08-30T19:01:19Z'
            - container: step-upload-sbom
              imageID: >-
                quay.io/redhat-appstudio/cosign@sha256:c883d6f8d39148f2cea71bff4622d196d89df3e510f36c140c097b932f0dd5d5
              name: upload-sbom
              terminated:
                containerID: >-
                  cri-o://bb40ca2fda83463cac9034e8c9f46279270d77df5f591f86cbac30eae8e9408f
                exitCode: 0
                finishedAt: '2023-08-30T19:01:29Z'
                message: >-
                  [{"key":"BASE_IMAGES_DIGESTS","value":"registry.access.redhat.com/ubi8/nodejs-18:latest@sha256:8a634a49c4e8d1753cddd3c047a4324d405fd2f6723f67b7ac274bb5fe72ff44\nregistry.access.redhat.com/ubi8/nodejs-18-minimal:latest@sha256:a1372e99afe18a64b9dbb70612452efe02e934db136fba00ba1da417ede408a0\n","type":1},{"key":"IMAGE_DIGEST","value":"sha256:9cc2777884b232b63f67c352ba96866edac8cf35fc14a87107bfcf64faaff7db","type":1},{"key":"IMAGE_URL","value":"quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002","type":1},{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]
                reason: Completed
                startedAt: '2023-08-30T19:01:28Z'
          taskResults:
            - name: JAVA_COMMUNITY_DEPENDENCIES
              type: string
              value: ''
            - name: BASE_IMAGES_DIGESTS
              type: string
              value: >
                registry.access.redhat.com/ubi8/nodejs-18:latest@sha256:8a634a49c4e8d1753cddd3c047a4324d405fd2f6723f67b7ac274bb5fe72ff44

                registry.access.redhat.com/ubi8/nodejs-18-minimal:latest@sha256:a1372e99afe18a64b9dbb70612452efe02e934db136fba00ba1da417ede408a0
            - name: IMAGE_DIGEST
              type: string
              value: >-
                sha256:9cc2777884b232b63f67c352ba96866edac8cf35fc14a87107bfcf64faaff7db
            - name: IMAGE_URL
              type: string
              value: >-
                quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
          taskSpec:
            description: >-
              Buildah task builds source code into a container image and pushes
              the image into container registry using buildah tool.

              In addition it generates a SBOM file, injects the SBOM file into
              final container image and pushes the SBOM file as separate image
              using cosign tool.

              When [Java dependency
              rebuild](https://redhat-appstudio.github.io/docs.stonesoup.io/Documentation/main/cli/proc_enabled_java_dependencies.html)
              is enabled it triggers rebuilds of Java artifacts.

              When prefetch-dependencies task was activated it is using its
              artifacts to run build in hermetic environment.
            params:
              - description: Reference of the image buildah will produce.
                name: IMAGE
                type: string
              - default: >-
                  registry.access.redhat.com/ubi9/buildah:9.0.0-19@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
                description: The location of the buildah builder image.
                name: BUILDER_IMAGE
                type: string
              - default: ./Dockerfile
                description: Path to the Dockerfile to build.
                name: DOCKERFILE
                type: string
              - default: .
                description: Path to the directory to use as context.
                name: CONTEXT
                type: string
              - default: 'true'
                description: >-
                  Verify the TLS on the registry endpoint (for push/pull to a
                  non-TLS registry)
                name: TLSVERIFY
                type: string
              - default: ''
                description: 'unused, should be removed in next task version'
                name: DOCKER_AUTH
                type: string
              - default: 'false'
                description: Determines if build will be executed without network access.
                name: HERMETIC
                type: string
              - default: ''
                description: >-
                  In case it is not empty, the prefetched content should be made
                  available to the build.
                name: PREFETCH_INPUT
                type: string
              - default: ''
                description: >-
                  Delete image tag after specified time. Empty means to keep the
                  image tag. Time values could be something like 1h, 2d, 3w for
                  hours, days, and weeks, respectively.
                name: IMAGE_EXPIRES_AFTER
                type: string
              - default: ''
                description: The image is built from this commit.
                name: COMMIT_SHA
                type: string
            results:
              - description: Digest of the image just built
                name: IMAGE_DIGEST
                type: string
              - description: Image repository where the built image was pushed
                name: IMAGE_URL
                type: string
              - description: Digests of the base images used for build
                name: BASE_IMAGES_DIGESTS
                type: string
              - description: The counting of Java components by publisher in JSON format
                name: SBOM_JAVA_COMPONENTS_COUNT
                type: string
              - description: >-
                  The Java dependencies that came from community sources such as
                  Maven central.
                name: JAVA_COMMUNITY_DEPENDENCIES
                type: string
            stepTemplate:
              env:
                - name: BUILDAH_FORMAT
                  value: oci
                - name: STORAGE_DRIVER
                  value: vfs
                - name: HERMETIC
                  value: 'false'
                - name: PREFETCH_INPUT
                - name: CONTEXT
                  value: .
                - name: DOCKERFILE
                  value: Dockerfile
                - name: IMAGE
                  value: >-
                    quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
                - name: TLSVERIFY
                  value: 'true'
                - name: IMAGE_EXPIRES_AFTER
              name: ''
              resources: {}
            steps:
              - env:
                  - name: COMMIT_SHA
                    value: cb64992eebf5c18900d283a9bca08b4ab5db2874
                image: 'quay.io/redhat-appstudio/buildah:v1.28'
                name: build
                resources:
                  limits:
                    cpu: '2'
                    memory: 4Gi
                  requests:
                    cpu: 250m
                    memory: 512Mi
                script: >
                  if [ -e "$CONTEXT/$DOCKERFILE" ]; then
                    dockerfile_path="$CONTEXT/$DOCKERFILE"
                  elif [ -e "$DOCKERFILE" ]; then
                    dockerfile_path="$DOCKERFILE"
                  elif echo "$DOCKERFILE" | grep -q "^https\?://"; then
                    echo "Fetch Dockerfile from $DOCKERFILE"
                    dockerfile_path=$(mktemp --suffix=-Dockerfile)
                    http_code=$(curl -s -L -w "%{http_code}" --output "$dockerfile_path" "$DOCKERFILE")
                    if [ $http_code != 200 ]; then
                      echo "No Dockerfile is fetched. Server responds $http_code"
                      exit 1
                    fi
                    http_code=$(curl -s -L -w "%{http_code}" --output "$dockerfile_path.dockerignore.tmp" "$DOCKERFILE.dockerignore")
                    if [ $http_code = 200 ]; then
                      echo "Fetched .dockerignore from $DOCKERFILE.dockerignore"
                      mv "$dockerfile_path.dockerignore.tmp" $CONTEXT/.dockerignore
                    fi
                  else
                    echo "Cannot find Dockerfile $DOCKERFILE"
                    exit 1
                  fi

                  if [ -n "$JVM_BUILD_WORKSPACE_ARTIFACT_CACHE_PORT_80_TCP_ADDR"
                  ] && grep -q '^\s*RUN \(./\)\?mvn' "$dockerfile_path"; then
                    sed -i -e "s|^\s*RUN \(\(./\)\?mvn\(.*\)\)|RUN echo \"<settings><mirrors><mirror><id>mirror.default</id><url>http://$JVM_BUILD_WORKSPACE_ARTIFACT_CACHE_PORT_80_TCP_ADDR/v1/cache/default/0/</url><mirrorOf>*</mirrorOf></mirror></mirrors></settings>\" > /tmp/settings.yaml; \1 -s /tmp/settings.yaml|g" "$dockerfile_path"
                    touch /var/lib/containers/java
                  fi


                  # Fixing group permission on /var/lib/containers

                  chown root:root /var/lib/containers


                  sed -i 's/^\s*short-name-mode\s*=\s*.*/short-name-mode =
                  "disabled"/' /etc/containers/registries.conf


                  # Setting new namespace to run buildah - 2^32-2

                  echo 'root:1:4294967294' | tee -a /etc/subuid >> /etc/subgid


                  if [ "${HERMETIC}" == "true" ]; then
                    BUILDAH_ARGS="--pull=never"
                    UNSHARE_ARGS="--net"
                    for image in $(grep -i '^\s*FROM' "$dockerfile_path" | sed 's/--platform=\S*//' | awk '{print $2}'); do
                      unshare -Ufp --keep-caps -r --map-users 1,1,65536 --map-groups 1,1,65536 -- buildah pull $image
                    done
                    echo "Build will be executed with network isolation"
                  fi


                  if [ -n "${PREFETCH_INPUT}" ]; then
                    mv cachi2 /tmp/
                    chmod -R go+rwX /tmp/cachi2
                    VOLUME_MOUNTS="--volume /tmp/cachi2:/cachi2"
                    sed -i 's|^\s*run |RUN . /cachi2/cachi2.env \&\& \\\n    |i' "$dockerfile_path"
                    echo "Prefetched content will be made available"
                  fi


                  LABELS=(
                    "--label" "build-date=$(date -u +'%Y-%m-%dT%H:%M:%S')"
                    "--label" "architecture=$(uname -m)"
                    "--label" "vcs-type=git"
                  )

                  [ -n "$COMMIT_SHA" ] && LABELS+=("--label"
                  "vcs-ref=$COMMIT_SHA")

                  [ -n "$IMAGE_EXPIRES_AFTER" ] && LABELS+=("--label"
                  "quay.expires-after=$IMAGE_EXPIRES_AFTER")


                  unshare -Uf $UNSHARE_ARGS --keep-caps -r --map-users 1,1,65536
                  --map-groups 1,1,65536 -- buildah build \
                    $VOLUME_MOUNTS \
                    $BUILDAH_ARGS \
                    ${LABELS[@]} \
                    --tls-verify=$TLSVERIFY --no-cache \
                    --ulimit nofile=4096:4096 \
                    -f "$dockerfile_path" -t $IMAGE $CONTEXT

                  container=$(buildah from --pull-never $IMAGE)

                  buildah mount $container | tee /workspace/container_path

                  echo $container > /workspace/container_name


                  # Save the SBOM produced by Cachi2 so it can be merged into
                  the final SBOM later

                  if [ -n "${PREFETCH_INPUT}" ]; then
                    cp /tmp/cachi2/output/bom.json ./sbom-cachi2.json
                  fi
                securityContext:
                  capabilities:
                    add:
                      - SETFCAP
                volumeMounts:
                  - mountPath: /var/lib/containers
                    name: varlibcontainers
                workingDir: /workspace/source
              - image: 'quay.io/redhat-appstudio/syft:v0.85.0'
                name: sbom-syft-generate
                resources: {}
                script: >
                  syft dir:/workspace/source
                  --file=/workspace/source/sbom-source.json
                  --output=cyclonedx-json

                  find $(cat /workspace/container_path) -xtype l -delete

                  syft dir:$(cat /workspace/container_path)
                  --file=/workspace/source/sbom-image.json
                  --output=cyclonedx-json
                volumeMounts:
                  - mountPath: /var/lib/containers
                    name: varlibcontainers
              - image: >-
                  quay.io/redhat-appstudio/hacbs-jvm-build-request-processor:1d417e6f1f3e68c6c537333b5759796eddae0afc
                name: analyse-dependencies-java-sbom
                resources: {}
                script: |
                  if [ -f /var/lib/containers/java ]; then
                    /opt/jboss/container/java/run/run-java.sh analyse-dependencies path $(cat /workspace/container_path) -s /workspace/source/sbom-image.json --task-run-name devfile-sample-nsj6m-build-container --publishers /tekton/results/SBOM_JAVA_COMPONENTS_COUNT
                    sed -i 's/^/ /' /tekton/results/SBOM_JAVA_COMPONENTS_COUNT # Workaround for SRVKP-2875
                  else
                    touch /tekton/results/JAVA_COMMUNITY_DEPENDENCIES
                  fi
                securityContext:
                  runAsUser: 0
                volumeMounts:
                  - mountPath: /var/lib/containers
                    name: varlibcontainers
              - image: 'registry.access.redhat.com/ubi9/python-39:1-133.1692772345'
                name: merge-syft-sboms
                resources: {}
                script: >
                  #!/bin/python3

                  import json


                  # load SBOMs

                  with open("./sbom-image.json") as f:
                    image_sbom = json.load(f)

                  with open("./sbom-source.json") as f:
                    source_sbom = json.load(f)

                  # fetch unique components from available SBOMs

                  def get_identifier(component):
                    return component["name"] + '@' + component.get("version", "")

                  existing_components = [get_identifier(component) for component
                  in image_sbom["components"]]


                  for component in source_sbom["components"]:
                    if get_identifier(component) not in existing_components:
                      image_sbom["components"].append(component)
                      existing_components.append(get_identifier(component))

                  image_sbom["components"].sort(key=lambda c: get_identifier(c))


                  # write the CycloneDX unified SBOM

                  with open("./sbom-cyclonedx.json", "w") as f:
                    json.dump(image_sbom, f, indent=4)
                securityContext:
                  runAsUser: 0
                workingDir: /workspace/source
              - image: >-
                  quay.io/redhat-appstudio/cachi2:0.3.0@sha256:46097f22b57e4d48a3fce96d931e08ccfe3a3e6421362d5f9353961279078eef
                name: merge-cachi2-sbom
                resources: {}
                script: |
                  if [ -n "${PREFETCH_INPUT}" ]; then
                    echo "Merging contents of sbom-cachi2.json into sbom-cyclonedx.json"
                    /src/utils/merge_syft_sbom.py sbom-cachi2.json sbom-cyclonedx.json > sbom-temp.json
                    mv sbom-temp.json sbom-cyclonedx.json
                  else
                    echo "Skipping step since no Cachi2 SBOM was produced"
                  fi
                securityContext:
                  runAsUser: 0
                workingDir: /workspace/source
              - image: 'registry.access.redhat.com/ubi9/python-39:1-133.1692772345'
                name: create-purl-sbom
                resources: {}
                script: >
                  #!/bin/python3

                  import json


                  with open("./sbom-cyclonedx.json") as f:
                    cyclonedx_sbom = json.load(f)

                  purls = [{"purl": component["purl"]} for component in
                  cyclonedx_sbom["components"] if "purl" in component]

                  purl_content = {"image_contents": {"dependencies": purls}}


                  with open("sbom-purl.json", "w") as output_file:
                    json.dump(purl_content, output_file, indent=4)
                securityContext:
                  runAsUser: 0
                workingDir: /workspace/source
              - image: >-
                  registry.access.redhat.com/ubi9/buildah:9.0.0-19@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
                name: inject-sbom-and-push
                resources: {}
                script: >
                  # Expose base image digests

                  buildah images --format '{{ .Name }}:{{ .Tag }}@{{ .Digest }}'
                  | grep -v $IMAGE > /tekton/results/BASE_IMAGES_DIGESTS


                  base_image_name=$(buildah inspect --format '{{ index
                  .ImageAnnotations "org.opencontainers.image.base.name"}}'
                  $IMAGE | cut -f1 -d'@')

                  base_image_digest=$(buildah inspect --format '{{ index
                  .ImageAnnotations "org.opencontainers.image.base.digest"}}'
                  $IMAGE)

                  container=$(buildah from --pull-never $IMAGE)

                  buildah copy $container sbom-cyclonedx.json sbom-purl.json
                  /root/buildinfo/content_manifests/

                  buildah config -a
                  org.opencontainers.image.base.name=${base_image_name} -a
                  org.opencontainers.image.base.digest=${base_image_digest}
                  $container

                  buildah commit $container $IMAGE


                  status=-1

                  max_run=5

                  sleep_sec=10

                  for run in $(seq 1 $max_run); do
                    status=0
                    [ "$run" -gt 1 ] && sleep $sleep_sec
                    echo "Pushing sbom image to registry"
                    buildah push \
                      --tls-verify=$TLSVERIFY \
                      --digestfile /workspace/source/image-digest $IMAGE \
                      docker://$IMAGE && break || status=$?
                  done

                  if [ "$status" -ne 0 ]; then
                      echo "Failed to push sbom image to registry after ${max_run} tries"
                      exit 1
                  fi


                  cat "/workspace/source"/image-digest | tee
                  /tekton/results/IMAGE_DIGEST

                  echo -n "$IMAGE" | tee /tekton/results/IMAGE_URL
                securityContext:
                  capabilities:
                    add:
                      - SETFCAP
                  runAsUser: 0
                volumeMounts:
                  - mountPath: /var/lib/containers
                    name: varlibcontainers
                workingDir: /workspace/source
              - args:
                  - attach
                  - sbom
                  - '--sbom'
                  - sbom-cyclonedx.json
                  - '--type'
                  - cyclonedx
                  - >-
                    quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
                image: 'quay.io/redhat-appstudio/cosign:v2.1.1'
                name: upload-sbom
                resources: {}
                workingDir: /workspace/source
            volumes:
              - emptyDir:
                  medium: Memory
                name: varlibcontainers
            workspaces:
              - description: Workspace containing the source code to build.
                name: source
          duration: 1m 1s
          reason: Succeeded
      - name: inspect-image
        params:
          - name: IMAGE_URL
            value: $(tasks.build-container.results.IMAGE_URL)
          - name: IMAGE_DIGEST
            value: $(tasks.build-container.results.IMAGE_DIGEST)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: inspect-image
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:inspect-image-0.1@sha256:bbc286f0a2ad94e671ceb9d0f1debd96f36b8c38c1147c5030957820b4125fc6
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'true'
            operator: in
            values:
              - 'false'
        workspaces:
          - name: source
            workspace: workspace
        status:
          reason: Skipped
      - name: deprecated-base-image-check
        params:
          - name: BASE_IMAGES_DIGESTS
            value: $(tasks.build-container.results.BASE_IMAGES_DIGESTS)
        taskRef:
          kind: Task
          params:
            - name: name
              value: deprecated-image-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:deprecated-image-check-0.2@sha256:58d16de95b4ca597f7f860fb85d6206e549910fa7a8d2a2cc229558f791ad329
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'true'
            operator: in
            values:
              - 'false'
        workspaces:
          - name: test-ws
            workspace: workspace
        status:
          reason: Skipped
      - name: clair-scan
        params:
          - name: image-digest
            value: $(tasks.build-container.results.IMAGE_DIGEST)
          - name: image-url
            value: $(tasks.build-container.results.IMAGE_URL)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: clair-scan
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:clair-scan-0.1@sha256:c5602d9d6dd797da98e98fde8471ea55a788c30f74f2192807910ce5436e9b66
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'true'
            operator: in
            values:
              - 'false'
        status:
          reason: Skipped
      - name: sast-snyk-check
        runAfter:
          - clone-repository
        taskRef:
          kind: Task
          params:
            - name: name
              value: sast-snyk-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:sast-snyk-check-0.1@sha256:9dcd450b454705b9fe22c5f8f7bb7305cebc3cb73e783b85e047f7e721994189
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'true'
            operator: in
            values:
              - 'false'
        workspaces:
          - name: workspace
            workspace: workspace
        status:
          reason: Skipped
      - name: clamav-scan
        params:
          - name: image-digest
            value: $(tasks.build-container.results.IMAGE_DIGEST)
          - name: image-url
            value: $(tasks.build-container.results.IMAGE_URL)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: clamav-scan
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:clamav-scan-0.1@sha256:cd4e301dd849cbdf7b8e38fd8f4915970b5b60174770df632a6b38ea93028d44
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'true'
            operator: in
            values:
              - 'false'
        status:
          reason: Skipped
      - name: sbom-json-check
        params:
          - name: IMAGE_URL
            value: $(tasks.build-container.results.IMAGE_URL)
          - name: IMAGE_DIGEST
            value: $(tasks.build-container.results.IMAGE_DIGEST)
        runAfter:
          - build-container
        taskRef:
          kind: Task
          params:
            - name: name
              value: sbom-json-check
            - name: bundle
              value: >-
                quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:sbom-json-check-0.1@sha256:397cb2fb20f413dec9653134231bec86edb80806a3441081fbf473677fc40917
            - name: kind
              value: task
          resolver: bundles
        when:
          - input: 'true'
            operator: in
            values:
              - 'false'
        status:
          reason: Skipped
    workspaces:
      - name: workspace
      - name: git-auth
        optional: true
  skippedTasks:
    - name: prefetch-dependencies
      reason: When Expressions evaluated to false
      whenExpressions:
        - input: 'false'
          operator: in
          values:
            - 'true'
    - name: inspect-image
      reason: When Expressions evaluated to false
      whenExpressions:
        - input: 'true'
          operator: in
          values:
            - 'false'
    - name: deprecated-base-image-check
      reason: When Expressions evaluated to false
      whenExpressions:
        - input: 'true'
          operator: in
          values:
            - 'false'
    - name: clair-scan
      reason: When Expressions evaluated to false
      whenExpressions:
        - input: 'true'
          operator: in
          values:
            - 'false'
    - name: sast-snyk-check
      reason: When Expressions evaluated to false
      whenExpressions:
        - input: 'true'
          operator: in
          values:
            - 'false'
    - name: clamav-scan
      reason: When Expressions evaluated to false
      whenExpressions:
        - input: 'true'
          operator: in
          values:
            - 'false'
    - name: sbom-json-check
      reason: When Expressions evaluated to false
      whenExpressions:
        - input: 'true'
          operator: in
          values:
            - 'false'
  startTime: '2023-08-30T19:00:02Z'
`

const tooBigNumTRInitYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipeline.tekton.dev/release: b8ad1b2
    tekton.dev/tags: 'appstudio, hacbs'
    results.tekton.dev/record: >-
      test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45/records/cfb31f93-a5de-4cf6-8b5b-14ee29c4e351
    results.tekton.dev/log: >-
      test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45/logs/08f293dc-fdcc-3ad8-a3d1-b2ee7dc10f9c
    build.appstudio.redhat.com/pipeline_name: docker-build
    build.appstudio.openshift.io/repo: >-
      https://github.com/nodeshift-starters/devfile-sample?rev=cb64992eebf5c18900d283a9bca08b4ab5db2874
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    results.tekton.dev/result: test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45
    build.appstudio.redhat.com/target_branch: main
    results.tekton.dev/childReadyForDeletion: 'true'
    build.appstudio.redhat.com/bundle: >-
      quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:docker-build-124dd76239d17a3bdb936ed26e403c9c9c4e947b
    build.appstudio.redhat.com/commit_sha: cb64992eebf5c18900d283a9bca08b4ab5db2874
  name: devfile-sample-nsj6m-init
  uid: cfb31f93-a5de-4cf6-8b5b-14ee29c4e351
  creationTimestamp: '2023-08-30T19:00:04Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:build.appstudio.redhat.com/pipeline_name': {}
            'f:build.appstudio.redhat.com/bundle': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            .: {}
            'f:pipeline.tekton.dev/release': {}
            'f:tekton.dev/tags': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:tekton.dev/pipelines.minVersion': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelines.openshift.io/used-by': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:appstudio.openshift.io/application': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:appstudio.openshift.io/component': {}
            'f:pipelines.openshift.io/strategy': {}
            'f:pipelines.openshift.io/runtime': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"579cdd9f-5dcc-436c-ab9b-eb770530ff45"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-30T19:00:04Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-30T19:00:09Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskResults': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:steps': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-30T19:00:09Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
            'f:results.tekton.dev/result': {}
      manager: watcher
      operation: Update
      time: '2023-08-30T19:00:29Z'
  namespace: test-rhtap-95-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: devfile-sample-nsj6m
      uid: 579cdd9f-5dcc-436c-ab9b-eb770530ff45
  finalizers:
    - chains.tekton.dev
  labels:
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: devfile-sample
    pipelines.openshift.io/runtime: generic
    pipelines.openshift.io/strategy: docker
    app.kubernetes.io/version: '0.1'
    tekton.dev/pipeline: docker-build
    pipelines.openshift.io/used-by: build-cloud
    app.kubernetes.io/managed-by: tekton-pipelines
    appstudio.openshift.io/application: load-app-7089021e-3585-438c-85e8-0a1ec8381a9f
    tekton.dev/task: init
    tekton.dev/pipelineTask: init
    pipelines.appstudio.openshift.io/type: build
    tekton.dev/pipelineRun: devfile-sample-nsj6m
spec:
  params:
    - name: image-url
      value: >-
        quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
    - name: rebuild
      value: 'false'
    - name: skip-checks
      value: 'true'
    - name: skip-optional
      value: 'true'
    - name: pipelinerun-name
      value: devfile-sample-nsj6m
    - name: pipelinerun-uid
      value: 579cdd9f-5dcc-436c-ab9b-eb770530ff45
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: init
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:init-0.1@sha256:26586a7ef08c3e86dfdaf0a5cc38dd3d70c4c02db1331b469caaed0a0f5b3d86
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
status:
  completionTime: '2023-08-30T19:00:09Z'
  conditions:
    - lastTransitionTime: '2023-08-30T19:00:09Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: devfile-sample-nsj6m-init-pod
  startTime: '2023-08-30T19:00:04Z'
  steps:
    - container: step-init
      imageID: >-
        registry.redhat.io/openshift4/ose-cli@sha256:73df37794ffff7de1101016c23dc623e4990810390ebdabcbbfa065214352c7c
      name: init
      terminated:
        containerID: >-
          cri-o://916f5f33bbb4819b281ba24ab276c20524c3ff33ad0fa86c213924255cadf59f
        exitCode: 0
        finishedAt: '2023-08-30T19:00:08Z'
        message: >-
          [{"key":"build","value":"true","type":1},{"key":"container-registry-secret","value":"unused\n","type":1}]
        reason: Completed
        startedAt: '2023-08-30T19:00:08Z'
  taskResults:
    - name: build
      type: string
      value: 'true'
    - name: container-registry-secret
      type: string
      value: |
        unused
  taskSpec:
    description: >-
      Initialize Pipeline Task, include flags for rebuild and auth. Generates
      image repository secret used by the PipelineRun.
    params:
      - description: Image URL for build by PipelineRun
        name: image-url
        type: string
      - default: 'false'
        description: Rebuild the image if exists
        name: rebuild
        type: string
      - default: 'false'
        description: Skip checks against built image
        name: skip-checks
        type: string
      - default: 'true'
        description: 'Skip optional checks, set false if you want to run optional checks'
        name: skip-optional
        type: string
      - default: ''
        description: 'unused, should be removed in next task version'
        name: pipelinerun-name
        type: string
      - default: ''
        description: 'unused, should be removed in next task version'
        name: pipelinerun-uid
        type: string
    results:
      - description: Defines if the image in param image-url should be built
        name: build
        type: string
      - description: 'unused, should be removed in next task version'
        name: container-registry-secret
        type: string
    steps:
      - env:
          - name: IMAGE_URL
            value: >-
              quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
          - name: REBUILD
            value: 'false'
          - name: SKIP_CHECKS
            value: 'true'
          - name: SKIP_OPTIONAL
            value: 'true'
        image: >-
          registry.redhat.io/openshift4/ose-cli:4.13@sha256:73df37794ffff7de1101016c23dc623e4990810390ebdabcbbfa065214352c7c
        name: init
        resources: {}
        script: >
          #!/bin/bash

          echo "Build Initialize: $IMAGE_URL"

          echo


          echo "Determine if Image Already Exists"

          # Build the image when image does not exists or rebuild is set to true

          if ! oc image info $IMAGE_URL &>/dev/null || [ "$REBUILD" == "true" ]
          || [ "$SKIP_CHECKS" == "false" ]; then
            echo -n "true" > /tekton/results/build
          else
            echo -n "false" > /tekton/results/build
          fi

          echo unused > /tekton/results/container-registry-secret
`

const tooBigNumTRCloneYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipeline.tekton.dev/release: b8ad1b2
    tekton.dev/tags: git
    results.tekton.dev/record: >-
      test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45/records/8355a9eb-c535-4330-bdd4-2a21a786ede5
    tekton.dev/categories: Git
    results.tekton.dev/log: >-
      test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45/logs/90081c05-9b81-30e4-91d1-b1ddee118de5
    build.appstudio.redhat.com/pipeline_name: docker-build
    build.appstudio.openshift.io/repo: >-
      https://github.com/nodeshift-starters/devfile-sample?rev=cb64992eebf5c18900d283a9bca08b4ab5db2874
    chains.tekton.dev/signed: 'true'
    tekton.dev/platforms: 'linux/amd64,linux/s390x,linux/ppc64le,linux/arm64'
    tekton.dev/pipelines.minVersion: 0.21.0
    results.tekton.dev/result: test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45
    build.appstudio.redhat.com/target_branch: main
    results.tekton.dev/childReadyForDeletion: 'true'
    build.appstudio.redhat.com/bundle: >-
      quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:docker-build-124dd76239d17a3bdb936ed26e403c9c9c4e947b
    tekton.dev/displayName: git clone
    build.appstudio.redhat.com/commit_sha: cb64992eebf5c18900d283a9bca08b4ab5db2874
  name: devfile-sample-nsj6m-clone-repository
  uid: 8355a9eb-c535-4330-bdd4-2a21a786ede5
  creationTimestamp: '2023-08-30T19:00:09Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:build.appstudio.redhat.com/pipeline_name': {}
            'f:tekton.dev/platforms': {}
            'f:build.appstudio.redhat.com/bundle': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            .: {}
            'f:pipeline.tekton.dev/release': {}
            'f:tekton.dev/displayName': {}
            'f:tekton.dev/tags': {}
            'f:tekton.dev/categories': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:tekton.dev/pipelines.minVersion': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelines.openshift.io/used-by': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:appstudio.openshift.io/application': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:appstudio.openshift.io/component': {}
            'f:pipelines.openshift.io/strategy': {}
            'f:pipelines.openshift.io/runtime': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"579cdd9f-5dcc-436c-ab9b-eb770530ff45"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
          'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-30T19:00:09Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-30T19:00:29Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskResults': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:steps': {}
            'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-30T19:00:29Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-30T19:00:37Z'
  namespace: test-rhtap-95-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: devfile-sample-nsj6m
      uid: 579cdd9f-5dcc-436c-ab9b-eb770530ff45
  finalizers:
    - chains.tekton.dev
  labels:
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: devfile-sample
    pipelines.openshift.io/runtime: generic
    pipelines.openshift.io/strategy: docker
    app.kubernetes.io/version: '0.1'
    tekton.dev/pipeline: docker-build
    pipelines.openshift.io/used-by: build-cloud
    app.kubernetes.io/managed-by: tekton-pipelines
    appstudio.openshift.io/application: load-app-7089021e-3585-438c-85e8-0a1ec8381a9f
    tekton.dev/task: git-clone
    tekton.dev/pipelineTask: clone-repository
    pipelines.appstudio.openshift.io/type: build
    tekton.dev/pipelineRun: devfile-sample-nsj6m
spec:
  params:
    - name: url
      value: 'https://github.com/nodeshift-starters/devfile-sample.git'
    - name: revision
      value: main
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: git-clone
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:git-clone-0.1@sha256:1f84973a21aabea38434b1f663abc4cb2d86565a9c7aae1f90decb43a8fa48eb
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
  workspaces:
    - name: output
      persistentVolumeClaim:
        claimName: pvc-a54b92c8ba
status:
  completionTime: '2023-08-30T19:00:29Z'
  conditions:
    - lastTransitionTime: '2023-08-30T19:00:29Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: devfile-sample-nsj6m-clone-repository-pod
  startTime: '2023-08-30T19:00:09Z'
  steps:
    - container: step-clone
      imageID: >-
        registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8@sha256:2fa0b06d52b04f377c696412e19307a9eff27383f81d87aae0b4f71672a1cd0b
      name: clone
      terminated:
        containerID: >-
          cri-o://2f18fd7be23d4e6ac18bd3066b7f4efbedb0d4527f29c321ccf50e5b3187ae8d
        exitCode: 0
        finishedAt: '2023-08-30T19:00:28Z'
        message: >-
          [{"key":"commit","value":"cb64992eebf5c18900d283a9bca08b4ab5db2874","type":1},{"key":"url","value":"https://github.com/nodeshift-starters/devfile-sample.git","type":1}]
        reason: Completed
        startedAt: '2023-08-30T19:00:27Z'
    - container: step-symlink-check
      imageID: >-
        registry.redhat.io/ubi9@sha256:089bd3b82a78ac45c0eed231bb58bfb43bfcd0560d9bba240fc6355502c92976
      name: symlink-check
      terminated:
        containerID: >-
          cri-o://1492a8338e20bcc7897d81af16d33e47ca26a51b45e1c238ff657610b48a243a
        exitCode: 0
        finishedAt: '2023-08-30T19:00:28Z'
        message: >-
          [{"key":"commit","value":"cb64992eebf5c18900d283a9bca08b4ab5db2874","type":1},{"key":"url","value":"https://github.com/nodeshift-starters/devfile-sample.git","type":1}]
        reason: Completed
        startedAt: '2023-08-30T19:00:28Z'
  taskResults:
    - name: commit
      type: string
      value: cb64992eebf5c18900d283a9bca08b4ab5db2874
    - name: url
      type: string
      value: 'https://github.com/nodeshift-starters/devfile-sample.git'
  taskSpec:
    description: >-
      The git-clone Task will clone a repo from the provided url into the output
      Workspace. By default the repo will be cloned into the root of your
      Workspace.
    params:
      - description: Repository URL to clone from.
        name: url
        type: string
      - default: ''
        description: 'Revision to checkout. (branch, tag, sha, ref, etc...)'
        name: revision
        type: string
      - default: ''
        description: Refspec to fetch before checking out revision.
        name: refspec
        type: string
      - default: 'true'
        description: Initialize and fetch git submodules.
        name: submodules
        type: string
      - default: '1'
        description: 'Perform a shallow clone, fetching only the most recent N commits.'
        name: depth
        type: string
      - default: 'true'
        description: >-
          Set the http.sslVerify global git config. Setting this to false is
          not advised unless you are sure that you trust your git remote.
        name: sslVerify
        type: string
      - default: ''
        description: Subdirectory inside the output Workspace to clone the repo into.
        name: subdirectory
        type: string
      - default: ''
        description: >-
          Define the directory patterns to match or exclude when performing a
          sparse checkout.
        name: sparseCheckoutDirectories
        type: string
      - default: 'true'
        description: >-
          Clean out the contents of the destination directory if it already
          exists before cloning.
        name: deleteExisting
        type: string
      - default: ''
        description: HTTP proxy server for non-SSL requests.
        name: httpProxy
        type: string
      - default: ''
        description: HTTPS proxy server for SSL requests.
        name: httpsProxy
        type: string
      - default: ''
        description: Opt out of proxying HTTP/HTTPS requests.
        name: noProxy
        type: string
      - default: 'true'
        description: Log the commands that are executed during git-clone's operation.
        name: verbose
        type: string
      - default: >-
          registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8:v1.8.2-8@sha256:a538c423e7a11aae6ae582a411fdb090936458075f99af4ce5add038bb6983e8
        description: The image providing the git-init binary that this Task runs.
        name: gitInitImage
        type: string
      - default: /tekton/home
        description: >
          Absolute path to the user's home directory. Set this explicitly if you
          are running the image as a non-root user or have overridden

          the gitInitImage param with an image containing custom user
          configuration.
        name: userHome
        type: string
      - default: 'true'
        description: >
          Check symlinks in the repo. If they're pointing outside of the repo,
          the build will fail.
        name: enableSymlinkCheck
        type: string
    results:
      - description: The precise commit SHA that was fetched by this Task.
        name: commit
        type: string
      - description: The precise URL that was fetched by this Task.
        name: url
        type: string
    steps:
      - env:
          - name: HOME
            value: /tekton/home
          - name: PARAM_URL
            value: 'https://github.com/nodeshift-starters/devfile-sample.git'
          - name: PARAM_REVISION
            value: main
          - name: PARAM_REFSPEC
          - name: PARAM_SUBMODULES
            value: 'true'
          - name: PARAM_DEPTH
            value: '1'
          - name: PARAM_SSL_VERIFY
            value: 'true'
          - name: PARAM_SUBDIRECTORY
          - name: PARAM_DELETE_EXISTING
            value: 'true'
          - name: PARAM_HTTP_PROXY
          - name: PARAM_HTTPS_PROXY
          - name: PARAM_NO_PROXY
          - name: PARAM_VERBOSE
            value: 'true'
          - name: PARAM_SPARSE_CHECKOUT_DIRECTORIES
          - name: PARAM_USER_HOME
            value: /tekton/home
          - name: WORKSPACE_OUTPUT_PATH
            value: /workspace/output
          - name: WORKSPACE_SSH_DIRECTORY_BOUND
            value: 'false'
          - name: WORKSPACE_SSH_DIRECTORY_PATH
          - name: WORKSPACE_BASIC_AUTH_DIRECTORY_BOUND
            value: 'false'
          - name: WORKSPACE_BASIC_AUTH_DIRECTORY_PATH
        image: >-
          registry.redhat.io/openshift-pipelines/pipelines-git-init-rhel8:v1.8.2-8@sha256:a538c423e7a11aae6ae582a411fdb090936458075f99af4ce5add038bb6983e8
        name: clone
        resources: {}
        script: >
          #!/usr/bin/env sh

          set -eu


          if [ "${PARAM_VERBOSE}" = "true" ] ; then
            set -x
          fi


          if [ "${WORKSPACE_BASIC_AUTH_DIRECTORY_BOUND}" = "true" ] ; then
            if [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.git-credentials" ] && [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.gitconfig" ]; then
              cp "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.git-credentials" "${PARAM_USER_HOME}/.git-credentials"
              cp "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/.gitconfig" "${PARAM_USER_HOME}/.gitconfig"
            # Compatibility with kubernetes.io/basic-auth secrets
            elif [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/username" ] && [ -f "${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/password" ]; then
              HOSTNAME=$(echo $PARAM_URL | awk -F/ '{print $3}')
              echo "https://$(cat ${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/username):$(cat ${WORKSPACE_BASIC_AUTH_DIRECTORY_PATH}/password)@$HOSTNAME" > "${PARAM_USER_HOME}/.git-credentials"
              echo -e "[credential \"https://$HOSTNAME\"]\n  helper = store" > "${PARAM_USER_HOME}/.gitconfig"
            else
              echo "Unknown basic-auth workspace format"
              exit 1
            fi
            chmod 400 "${PARAM_USER_HOME}/.git-credentials"
            chmod 400 "${PARAM_USER_HOME}/.gitconfig"
          fi


          if [ "${WORKSPACE_SSH_DIRECTORY_BOUND}" = "true" ] ; then
            cp -R "${WORKSPACE_SSH_DIRECTORY_PATH}" "${PARAM_USER_HOME}"/.ssh
            chmod 700 "${PARAM_USER_HOME}"/.ssh
            chmod -R 400 "${PARAM_USER_HOME}"/.ssh/*
          fi


          CHECKOUT_DIR="${WORKSPACE_OUTPUT_PATH}/${PARAM_SUBDIRECTORY}"


          cleandir() {
            # Delete any existing contents of the repo directory if it exists.
            #
            # We don't just "rm -rf ${CHECKOUT_DIR}" because ${CHECKOUT_DIR} might be "/"
            # or the root of a mounted volume.
            if [ -d "${CHECKOUT_DIR}" ] ; then
              # Delete non-hidden files and directories
              rm -rf "${CHECKOUT_DIR:?}"/*
              # Delete files and directories starting with . but excluding ..
              rm -rf "${CHECKOUT_DIR}"/.[!.]*
              # Delete files and directories starting with .. plus any other character
              rm -rf "${CHECKOUT_DIR}"/..?*
            fi
          }


          if [ "${PARAM_DELETE_EXISTING}" = "true" ] ; then
            cleandir
          fi


          test -z "${PARAM_HTTP_PROXY}" || export
          HTTP_PROXY="${PARAM_HTTP_PROXY}"

          test -z "${PARAM_HTTPS_PROXY}" || export
          HTTPS_PROXY="${PARAM_HTTPS_PROXY}"

          test -z "${PARAM_NO_PROXY}" || export NO_PROXY="${PARAM_NO_PROXY}"


          /ko-app/git-init \
            -url="${PARAM_URL}" \
            -revision="${PARAM_REVISION}" \
            -refspec="${PARAM_REFSPEC}" \
            -path="${CHECKOUT_DIR}" \
            -sslVerify="${PARAM_SSL_VERIFY}" \
            -submodules="${PARAM_SUBMODULES}" \
            -depth="${PARAM_DEPTH}" \
            -sparseCheckoutDirectories="${PARAM_SPARSE_CHECKOUT_DIRECTORIES}"
          cd "${CHECKOUT_DIR}"

          RESULT_SHA="$(git rev-parse HEAD)"

          EXIT_CODE="$?"

          if [ "${EXIT_CODE}" != 0 ] ; then
            exit "${EXIT_CODE}"
          fi

          printf "%s" "${RESULT_SHA}" > "/tekton/results/commit"

          printf "%s" "${PARAM_URL}" > "/tekton/results/url"
        securityContext:
          runAsUser: 0
      - env:
          - name: PARAM_ENABLE_SYMLINK_CHECK
            value: 'true'
          - name: PARAM_SUBDIRECTORY
          - name: WORKSPACE_OUTPUT_PATH
            value: /workspace/output
        image: 'registry.redhat.io/ubi9:9.2-696'
        name: symlink-check
        resources: {}
        script: |
          #!/usr/bin/env bash
          set -euo pipefail

          CHECKOUT_DIR="${WORKSPACE_OUTPUT_PATH}/${PARAM_SUBDIRECTORY}"
          check_symlinks() {
            FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO=false
            while read symlink
            do
              target=$(readlink -f "$symlink")
              if ! [[ "$target" =~ ^$CHECKOUT_DIR ]]; then
                echo "The cloned repository contains symlink pointing outside of the cloned repository: $symlink"
                FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO=true
              fi
            done < <(find $CHECKOUT_DIR -type l -print)
            if [ "$FOUND_SYMLINK_POINTING_OUTSIDE_OF_REPO" = true ] ; then
              return 1
            fi
          }

          if [ "${PARAM_ENABLE_SYMLINK_CHECK}" = "true" ] ; then
            echo "Running symlink check"
            check_symlinks
          fi
    workspaces:
      - description: The git repo will be cloned onto the volume backing this Workspace.
        name: output
      - description: >
          A .ssh directory with private key, known_hosts, config, etc. Copied to

          the user's home before git commands are executed. Used to authenticate

          with the git remote when performing the clone. Binding a Secret to
          this

          Workspace is strongly recommended over other volume types.
        name: ssh-directory
        optional: true
      - description: >
          A Workspace containing a .gitconfig and .git-credentials file or
          username and password.

          These will be copied to the user's home before any git commands are
          run. Any

          other files in this Workspace are ignored. It is strongly recommended

          to use ssh-directory over basic-auth whenever possible and to bind a

          Secret to this Workspace over other volume types.
        name: basic-auth
        optional: true
`

const tooBigNumTRBuildYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipeline.tekton.dev/release: b8ad1b2
    tekton.dev/tags: 'image-build, appstudio, hacbs'
    results.tekton.dev/record: >-
      test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45/records/216501bb-80fa-4e91-965e-db097dbc477e
    results.tekton.dev/log: >-
      test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45/logs/14b06370-8d0b-3cee-8fbb-279acfb72b03
    build.appstudio.redhat.com/pipeline_name: docker-build
    build.appstudio.openshift.io/repo: >-
      https://github.com/nodeshift-starters/devfile-sample?rev=cb64992eebf5c18900d283a9bca08b4ab5db2874
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    results.tekton.dev/result: test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45
    build.appstudio.redhat.com/target_branch: main
    results.tekton.dev/childReadyForDeletion: 'true'
    build.appstudio.redhat.com/bundle: >-
      quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:docker-build-124dd76239d17a3bdb936ed26e403c9c9c4e947b
    build.appstudio.redhat.com/commit_sha: cb64992eebf5c18900d283a9bca08b4ab5db2874
  name: devfile-sample-nsj6m-build-container
  uid: 216501bb-80fa-4e91-965e-db097dbc477e
  creationTimestamp: '2023-08-30T19:00:29Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:build.appstudio.redhat.com/pipeline_name': {}
            'f:build.appstudio.redhat.com/bundle': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            .: {}
            'f:pipeline.tekton.dev/release': {}
            'f:tekton.dev/tags': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:tekton.dev/pipelines.minVersion': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelines.openshift.io/used-by': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:appstudio.openshift.io/application': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:build.appstudio.redhat.com/build_type': {}
            'f:appstudio.openshift.io/component': {}
            'f:pipelines.openshift.io/strategy': {}
            'f:pipelines.openshift.io/runtime': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"579cdd9f-5dcc-436c-ab9b-eb770530ff45"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
          'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-30T19:00:29Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskResults': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:results': {}
            'f:stepTemplate':
              .: {}
              'f:env': {}
              'f:name': {}
              'f:resources': {}
            'f:steps': {}
            'f:volumes': {}
            'f:workspaces': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-30T19:01:30Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-30T19:01:31Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-30T19:02:07Z'
  namespace: test-rhtap-95-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: devfile-sample-nsj6m
      uid: 579cdd9f-5dcc-436c-ab9b-eb770530ff45
  finalizers:
    - chains.tekton.dev
  labels:
    tekton.dev/memberOf: tasks
    appstudio.openshift.io/component: devfile-sample
    pipelines.openshift.io/runtime: generic
    build.appstudio.redhat.com/build_type: docker
    pipelines.openshift.io/strategy: docker
    app.kubernetes.io/version: '0.1'
    tekton.dev/pipeline: docker-build
    pipelines.openshift.io/used-by: build-cloud
    app.kubernetes.io/managed-by: tekton-pipelines
    appstudio.openshift.io/application: load-app-7089021e-3585-438c-85e8-0a1ec8381a9f
    tekton.dev/task: buildah
    tekton.dev/pipelineTask: build-container
    pipelines.appstudio.openshift.io/type: build
    tekton.dev/pipelineRun: devfile-sample-nsj6m
spec:
  params:
    - name: IMAGE
      value: >-
        quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
    - name: DOCKERFILE
      value: Dockerfile
    - name: CONTEXT
      value: .
    - name: HERMETIC
      value: 'false'
    - name: PREFETCH_INPUT
      value: ''
    - name: IMAGE_EXPIRES_AFTER
      value: ''
    - name: COMMIT_SHA
      value: cb64992eebf5c18900d283a9bca08b4ab5db2874
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: buildah
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:buildah-0.1@sha256:bfb5edabab8128e24608df608448b9392fd0a2b61ac05a53e83aa60d8929b73f
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
  workspaces:
    - name: source
      persistentVolumeClaim:
        claimName: pvc-a54b92c8ba
status:
  completionTime: '2023-08-30T19:01:30Z'
  conditions:
    - lastTransitionTime: '2023-08-30T19:01:30Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: devfile-sample-nsj6m-build-container-pod
  startTime: '2023-08-30T19:00:29Z'
  steps:
    - container: step-build
      imageID: >-
        quay.io/redhat-appstudio/buildah@sha256:381e9bfedd59701477621da93892106873a6951b196105d3d2d85c3f6d7b569b
      name: build
      terminated:
        containerID: >-
          cri-o://24d9f9c59f1c0fe114d05b1c29ac7209c3d3482716cc740f4783c8903c4721fb
        exitCode: 0
        finishedAt: '2023-08-30T19:01:16Z'
        reason: Completed
        startedAt: '2023-08-30T19:00:50Z'
    - container: step-sbom-syft-generate
      imageID: >-
        quay.io/redhat-appstudio/syft@sha256:244a17ce220a0b7a54c862c4fe3b72ce92799910c5eff8e94ac2f121fa5b4a53
      name: sbom-syft-generate
      terminated:
        containerID: >-
          cri-o://cb51f07d379e15f8a8755bfb84428ff60de6989f0a0c2a1de3fa07927b09ec2c
        exitCode: 0
        finishedAt: '2023-08-30T19:01:18Z'
        reason: Completed
        startedAt: '2023-08-30T19:01:16Z'
    - container: step-analyse-dependencies-java-sbom
      imageID: >-
        quay.io/redhat-appstudio/hacbs-jvm-build-request-processor@sha256:b198cf4b33dab59ce8ac25afd4e1001390db29ca2dec83dc8a1e21b0359ce743
      name: analyse-dependencies-java-sbom
      terminated:
        containerID: >-
          cri-o://b7448e62a57890749446e24a4f7f51ea2fa6333d9616c1193f4f92d9728e8fe6
        exitCode: 0
        finishedAt: '2023-08-30T19:01:18Z'
        message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
        reason: Completed
        startedAt: '2023-08-30T19:01:18Z'
    - container: step-merge-syft-sboms
      imageID: >-
        registry.access.redhat.com/ubi9/python-39@sha256:562b4f9cd4e5abeba54c56fd2290096a71f6e10aa8c1c18f43c9d6962c2d4d41
      name: merge-syft-sboms
      terminated:
        containerID: >-
          cri-o://7082cb7ddc3c629dca86aac4b5ae553eeb09960402ff503b49753b4f8c1ef7a1
        exitCode: 0
        finishedAt: '2023-08-30T19:01:19Z'
        message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
        reason: Completed
        startedAt: '2023-08-30T19:01:18Z'
    - container: step-merge-cachi2-sbom
      imageID: >-
        quay.io/redhat-appstudio/cachi2@sha256:46097f22b57e4d48a3fce96d931e08ccfe3a3e6421362d5f9353961279078eef
      name: merge-cachi2-sbom
      terminated:
        containerID: >-
          cri-o://543053c4be3079d0d266d494b520852f15d3329a5b2e5ab36f12ffb138eac2e5
        exitCode: 0
        finishedAt: '2023-08-30T19:01:19Z'
        message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
        reason: Completed
        startedAt: '2023-08-30T19:01:19Z'
    - container: step-create-purl-sbom
      imageID: >-
        registry.access.redhat.com/ubi9/python-39@sha256:562b4f9cd4e5abeba54c56fd2290096a71f6e10aa8c1c18f43c9d6962c2d4d41
      name: create-purl-sbom
      terminated:
        containerID: >-
          cri-o://7e41fead679f12a2291dc1c1d59c4fff229cfd7f17cebea7381aa9dd2c6472a5
        exitCode: 0
        finishedAt: '2023-08-30T19:01:19Z'
        message: '[{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]'
        reason: Completed
        startedAt: '2023-08-30T19:01:19Z'
    - container: step-inject-sbom-and-push
      imageID: >-
        registry.access.redhat.com/ubi9/buildah@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
      name: inject-sbom-and-push
      terminated:
        containerID: >-
          cri-o://f17b5f541734bb1c60ba25aa733925f009345ba2b47d77ff25b5c8d28f08802b
        exitCode: 0
        finishedAt: '2023-08-30T19:01:27Z'
        message: >-
          [{"key":"BASE_IMAGES_DIGESTS","value":"registry.access.redhat.com/ubi8/nodejs-18:latest@sha256:8a634a49c4e8d1753cddd3c047a4324d405fd2f6723f67b7ac274bb5fe72ff44\nregistry.access.redhat.com/ubi8/nodejs-18-minimal:latest@sha256:a1372e99afe18a64b9dbb70612452efe02e934db136fba00ba1da417ede408a0\n","type":1},{"key":"IMAGE_DIGEST","value":"sha256:9cc2777884b232b63f67c352ba96866edac8cf35fc14a87107bfcf64faaff7db","type":1},{"key":"IMAGE_URL","value":"quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002","type":1},{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]
        reason: Completed
        startedAt: '2023-08-30T19:01:19Z'
    - container: step-upload-sbom
      imageID: >-
        quay.io/redhat-appstudio/cosign@sha256:c883d6f8d39148f2cea71bff4622d196d89df3e510f36c140c097b932f0dd5d5
      name: upload-sbom
      terminated:
        containerID: >-
          cri-o://bb40ca2fda83463cac9034e8c9f46279270d77df5f591f86cbac30eae8e9408f
        exitCode: 0
        finishedAt: '2023-08-30T19:01:29Z'
        message: >-
          [{"key":"BASE_IMAGES_DIGESTS","value":"registry.access.redhat.com/ubi8/nodejs-18:latest@sha256:8a634a49c4e8d1753cddd3c047a4324d405fd2f6723f67b7ac274bb5fe72ff44\nregistry.access.redhat.com/ubi8/nodejs-18-minimal:latest@sha256:a1372e99afe18a64b9dbb70612452efe02e934db136fba00ba1da417ede408a0\n","type":1},{"key":"IMAGE_DIGEST","value":"sha256:9cc2777884b232b63f67c352ba96866edac8cf35fc14a87107bfcf64faaff7db","type":1},{"key":"IMAGE_URL","value":"quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002","type":1},{"key":"JAVA_COMMUNITY_DEPENDENCIES","value":"","type":1}]
        reason: Completed
        startedAt: '2023-08-30T19:01:28Z'
  taskResults:
    - name: JAVA_COMMUNITY_DEPENDENCIES
      type: string
      value: ''
    - name: BASE_IMAGES_DIGESTS
      type: string
      value: >
        registry.access.redhat.com/ubi8/nodejs-18:latest@sha256:8a634a49c4e8d1753cddd3c047a4324d405fd2f6723f67b7ac274bb5fe72ff44

        registry.access.redhat.com/ubi8/nodejs-18-minimal:latest@sha256:a1372e99afe18a64b9dbb70612452efe02e934db136fba00ba1da417ede408a0
    - name: IMAGE_DIGEST
      type: string
      value: 'sha256:9cc2777884b232b63f67c352ba96866edac8cf35fc14a87107bfcf64faaff7db'
    - name: IMAGE_URL
      type: string
      value: >-
        quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
  taskSpec:
    description: >-
      Buildah task builds source code into a container image and pushes the
      image into container registry using buildah tool.

      In addition it generates a SBOM file, injects the SBOM file into final
      container image and pushes the SBOM file as separate image using cosign
      tool.

      When [Java dependency
      rebuild](https://redhat-appstudio.github.io/docs.stonesoup.io/Documentation/main/cli/proc_enabled_java_dependencies.html)
      is enabled it triggers rebuilds of Java artifacts.

      When prefetch-dependencies task was activated it is using its artifacts to
      run build in hermetic environment.
    params:
      - description: Reference of the image buildah will produce.
        name: IMAGE
        type: string
      - default: >-
          registry.access.redhat.com/ubi9/buildah:9.0.0-19@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
        description: The location of the buildah builder image.
        name: BUILDER_IMAGE
        type: string
      - default: ./Dockerfile
        description: Path to the Dockerfile to build.
        name: DOCKERFILE
        type: string
      - default: .
        description: Path to the directory to use as context.
        name: CONTEXT
        type: string
      - default: 'true'
        description: >-
          Verify the TLS on the registry endpoint (for push/pull to a non-TLS
          registry)
        name: TLSVERIFY
        type: string
      - default: ''
        description: 'unused, should be removed in next task version'
        name: DOCKER_AUTH
        type: string
      - default: 'false'
        description: Determines if build will be executed without network access.
        name: HERMETIC
        type: string
      - default: ''
        description: >-
          In case it is not empty, the prefetched content should be made
          available to the build.
        name: PREFETCH_INPUT
        type: string
      - default: ''
        description: >-
          Delete image tag after specified time. Empty means to keep the image
          tag. Time values could be something like 1h, 2d, 3w for hours, days,
          and weeks, respectively.
        name: IMAGE_EXPIRES_AFTER
        type: string
      - default: ''
        description: The image is built from this commit.
        name: COMMIT_SHA
        type: string
    results:
      - description: Digest of the image just built
        name: IMAGE_DIGEST
        type: string
      - description: Image repository where the built image was pushed
        name: IMAGE_URL
        type: string
      - description: Digests of the base images used for build
        name: BASE_IMAGES_DIGESTS
        type: string
      - description: The counting of Java components by publisher in JSON format
        name: SBOM_JAVA_COMPONENTS_COUNT
        type: string
      - description: >-
          The Java dependencies that came from community sources such as Maven
          central.
        name: JAVA_COMMUNITY_DEPENDENCIES
        type: string
    stepTemplate:
      env:
        - name: BUILDAH_FORMAT
          value: oci
        - name: STORAGE_DRIVER
          value: vfs
        - name: HERMETIC
          value: 'false'
        - name: PREFETCH_INPUT
        - name: CONTEXT
          value: .
        - name: DOCKERFILE
          value: Dockerfile
        - name: IMAGE
          value: >-
            quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
        - name: TLSVERIFY
          value: 'true'
        - name: IMAGE_EXPIRES_AFTER
      name: ''
      resources: {}
    steps:
      - env:
          - name: COMMIT_SHA
            value: cb64992eebf5c18900d283a9bca08b4ab5db2874
        image: 'quay.io/redhat-appstudio/buildah:v1.28'
        name: build
        resources:
          limits:
            cpu: '2'
            memory: 4Gi
          requests:
            cpu: 250m
            memory: 512Mi
        script: >
          if [ -e "$CONTEXT/$DOCKERFILE" ]; then
            dockerfile_path="$CONTEXT/$DOCKERFILE"
          elif [ -e "$DOCKERFILE" ]; then
            dockerfile_path="$DOCKERFILE"
          elif echo "$DOCKERFILE" | grep -q "^https\?://"; then
            echo "Fetch Dockerfile from $DOCKERFILE"
            dockerfile_path=$(mktemp --suffix=-Dockerfile)
            http_code=$(curl -s -L -w "%{http_code}" --output "$dockerfile_path" "$DOCKERFILE")
            if [ $http_code != 200 ]; then
              echo "No Dockerfile is fetched. Server responds $http_code"
              exit 1
            fi
            http_code=$(curl -s -L -w "%{http_code}" --output "$dockerfile_path.dockerignore.tmp" "$DOCKERFILE.dockerignore")
            if [ $http_code = 200 ]; then
              echo "Fetched .dockerignore from $DOCKERFILE.dockerignore"
              mv "$dockerfile_path.dockerignore.tmp" $CONTEXT/.dockerignore
            fi
          else
            echo "Cannot find Dockerfile $DOCKERFILE"
            exit 1
          fi

          if [ -n "$JVM_BUILD_WORKSPACE_ARTIFACT_CACHE_PORT_80_TCP_ADDR" ] &&
          grep -q '^\s*RUN \(./\)\?mvn' "$dockerfile_path"; then
            sed -i -e "s|^\s*RUN \(\(./\)\?mvn\(.*\)\)|RUN echo \"<settings><mirrors><mirror><id>mirror.default</id><url>http://$JVM_BUILD_WORKSPACE_ARTIFACT_CACHE_PORT_80_TCP_ADDR/v1/cache/default/0/</url><mirrorOf>*</mirrorOf></mirror></mirrors></settings>\" > /tmp/settings.yaml; \1 -s /tmp/settings.yaml|g" "$dockerfile_path"
            touch /var/lib/containers/java
          fi


          # Fixing group permission on /var/lib/containers

          chown root:root /var/lib/containers


          sed -i 's/^\s*short-name-mode\s*=\s*.*/short-name-mode = "disabled"/'
          /etc/containers/registries.conf


          # Setting new namespace to run buildah - 2^32-2

          echo 'root:1:4294967294' | tee -a /etc/subuid >> /etc/subgid


          if [ "${HERMETIC}" == "true" ]; then
            BUILDAH_ARGS="--pull=never"
            UNSHARE_ARGS="--net"
            for image in $(grep -i '^\s*FROM' "$dockerfile_path" | sed 's/--platform=\S*//' | awk '{print $2}'); do
              unshare -Ufp --keep-caps -r --map-users 1,1,65536 --map-groups 1,1,65536 -- buildah pull $image
            done
            echo "Build will be executed with network isolation"
          fi


          if [ -n "${PREFETCH_INPUT}" ]; then
            mv cachi2 /tmp/
            chmod -R go+rwX /tmp/cachi2
            VOLUME_MOUNTS="--volume /tmp/cachi2:/cachi2"
            sed -i 's|^\s*run |RUN . /cachi2/cachi2.env \&\& \\\n    |i' "$dockerfile_path"
            echo "Prefetched content will be made available"
          fi


          LABELS=(
            "--label" "build-date=$(date -u +'%Y-%m-%dT%H:%M:%S')"
            "--label" "architecture=$(uname -m)"
            "--label" "vcs-type=git"
          )

          [ -n "$COMMIT_SHA" ] && LABELS+=("--label" "vcs-ref=$COMMIT_SHA")

          [ -n "$IMAGE_EXPIRES_AFTER" ] && LABELS+=("--label"
          "quay.expires-after=$IMAGE_EXPIRES_AFTER")


          unshare -Uf $UNSHARE_ARGS --keep-caps -r --map-users 1,1,65536
          --map-groups 1,1,65536 -- buildah build \
            $VOLUME_MOUNTS \
            $BUILDAH_ARGS \
            ${LABELS[@]} \
            --tls-verify=$TLSVERIFY --no-cache \
            --ulimit nofile=4096:4096 \
            -f "$dockerfile_path" -t $IMAGE $CONTEXT

          container=$(buildah from --pull-never $IMAGE)

          buildah mount $container | tee /workspace/container_path

          echo $container > /workspace/container_name


          # Save the SBOM produced by Cachi2 so it can be merged into the final
          SBOM later

          if [ -n "${PREFETCH_INPUT}" ]; then
            cp /tmp/cachi2/output/bom.json ./sbom-cachi2.json
          fi
        securityContext:
          capabilities:
            add:
              - SETFCAP
        volumeMounts:
          - mountPath: /var/lib/containers
            name: varlibcontainers
        workingDir: /workspace/source
      - image: 'quay.io/redhat-appstudio/syft:v0.85.0'
        name: sbom-syft-generate
        resources: {}
        script: >
          syft dir:/workspace/source --file=/workspace/source/sbom-source.json
          --output=cyclonedx-json

          find $(cat /workspace/container_path) -xtype l -delete

          syft dir:$(cat /workspace/container_path)
          --file=/workspace/source/sbom-image.json --output=cyclonedx-json
        volumeMounts:
          - mountPath: /var/lib/containers
            name: varlibcontainers
      - image: >-
          quay.io/redhat-appstudio/hacbs-jvm-build-request-processor:1d417e6f1f3e68c6c537333b5759796eddae0afc
        name: analyse-dependencies-java-sbom
        resources: {}
        script: |
          if [ -f /var/lib/containers/java ]; then
            /opt/jboss/container/java/run/run-java.sh analyse-dependencies path $(cat /workspace/container_path) -s /workspace/source/sbom-image.json --task-run-name devfile-sample-nsj6m-build-container --publishers /tekton/results/SBOM_JAVA_COMPONENTS_COUNT
            sed -i 's/^/ /' /tekton/results/SBOM_JAVA_COMPONENTS_COUNT # Workaround for SRVKP-2875
          else
            touch /tekton/results/JAVA_COMMUNITY_DEPENDENCIES
          fi
        securityContext:
          runAsUser: 0
        volumeMounts:
          - mountPath: /var/lib/containers
            name: varlibcontainers
      - image: 'registry.access.redhat.com/ubi9/python-39:1-133.1692772345'
        name: merge-syft-sboms
        resources: {}
        script: >
          #!/bin/python3

          import json


          # load SBOMs

          with open("./sbom-image.json") as f:
            image_sbom = json.load(f)

          with open("./sbom-source.json") as f:
            source_sbom = json.load(f)

          # fetch unique components from available SBOMs

          def get_identifier(component):
            return component["name"] + '@' + component.get("version", "")

          existing_components = [get_identifier(component) for component in
          image_sbom["components"]]


          for component in source_sbom["components"]:
            if get_identifier(component) not in existing_components:
              image_sbom["components"].append(component)
              existing_components.append(get_identifier(component))

          image_sbom["components"].sort(key=lambda c: get_identifier(c))


          # write the CycloneDX unified SBOM

          with open("./sbom-cyclonedx.json", "w") as f:
            json.dump(image_sbom, f, indent=4)
        securityContext:
          runAsUser: 0
        workingDir: /workspace/source
      - image: >-
          quay.io/redhat-appstudio/cachi2:0.3.0@sha256:46097f22b57e4d48a3fce96d931e08ccfe3a3e6421362d5f9353961279078eef
        name: merge-cachi2-sbom
        resources: {}
        script: |
          if [ -n "${PREFETCH_INPUT}" ]; then
            echo "Merging contents of sbom-cachi2.json into sbom-cyclonedx.json"
            /src/utils/merge_syft_sbom.py sbom-cachi2.json sbom-cyclonedx.json > sbom-temp.json
            mv sbom-temp.json sbom-cyclonedx.json
          else
            echo "Skipping step since no Cachi2 SBOM was produced"
          fi
        securityContext:
          runAsUser: 0
        workingDir: /workspace/source
      - image: 'registry.access.redhat.com/ubi9/python-39:1-133.1692772345'
        name: create-purl-sbom
        resources: {}
        script: >
          #!/bin/python3

          import json


          with open("./sbom-cyclonedx.json") as f:
            cyclonedx_sbom = json.load(f)

          purls = [{"purl": component["purl"]} for component in
          cyclonedx_sbom["components"] if "purl" in component]

          purl_content = {"image_contents": {"dependencies": purls}}


          with open("sbom-purl.json", "w") as output_file:
            json.dump(purl_content, output_file, indent=4)
        securityContext:
          runAsUser: 0
        workingDir: /workspace/source
      - image: >-
          registry.access.redhat.com/ubi9/buildah:9.0.0-19@sha256:c8b1d312815452964885680fc5bc8d99b3bfe9b6961228c71a09c72ca8e915eb
        name: inject-sbom-and-push
        resources: {}
        script: >
          # Expose base image digests

          buildah images --format '{{ .Name }}:{{ .Tag }}@{{ .Digest }}' | grep
          -v $IMAGE > /tekton/results/BASE_IMAGES_DIGESTS


          base_image_name=$(buildah inspect --format '{{ index .ImageAnnotations
          "org.opencontainers.image.base.name"}}' $IMAGE | cut -f1 -d'@')

          base_image_digest=$(buildah inspect --format '{{ index
          .ImageAnnotations "org.opencontainers.image.base.digest"}}' $IMAGE)

          container=$(buildah from --pull-never $IMAGE)

          buildah copy $container sbom-cyclonedx.json sbom-purl.json
          /root/buildinfo/content_manifests/

          buildah config -a
          org.opencontainers.image.base.name=${base_image_name} -a
          org.opencontainers.image.base.digest=${base_image_digest} $container

          buildah commit $container $IMAGE


          status=-1

          max_run=5

          sleep_sec=10

          for run in $(seq 1 $max_run); do
            status=0
            [ "$run" -gt 1 ] && sleep $sleep_sec
            echo "Pushing sbom image to registry"
            buildah push \
              --tls-verify=$TLSVERIFY \
              --digestfile /workspace/source/image-digest $IMAGE \
              docker://$IMAGE && break || status=$?
          done

          if [ "$status" -ne 0 ]; then
              echo "Failed to push sbom image to registry after ${max_run} tries"
              exit 1
          fi


          cat "/workspace/source"/image-digest | tee
          /tekton/results/IMAGE_DIGEST

          echo -n "$IMAGE" | tee /tekton/results/IMAGE_URL
        securityContext:
          capabilities:
            add:
              - SETFCAP
          runAsUser: 0
        volumeMounts:
          - mountPath: /var/lib/containers
            name: varlibcontainers
        workingDir: /workspace/source
      - args:
          - attach
          - sbom
          - '--sbom'
          - sbom-cyclonedx.json
          - '--type'
          - cyclonedx
          - >-
            quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
        image: 'quay.io/redhat-appstudio/cosign:v2.1.1'
        name: upload-sbom
        resources: {}
        workingDir: /workspace/source
    volumes:
      - emptyDir:
          medium: Memory
        name: varlibcontainers
    workspaces:
      - description: Workspace containing the source code to build.
        name: source
`

const tooBigNumTRSbomYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipeline.tekton.dev/release: b8ad1b2
    tekton.dev/tags: 'appstudio, hacbs'
    results.tekton.dev/record: >-
      test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45/records/82872a3b-5596-45f2-84ac-d499953ad295
    results.tekton.dev/log: >-
      test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45/logs/7ef4eb0b-de57-3ed5-8c16-acfe5a2bfdd1
    build.appstudio.redhat.com/pipeline_name: docker-build
    build.appstudio.openshift.io/repo: >-
      https://github.com/nodeshift-starters/devfile-sample?rev=cb64992eebf5c18900d283a9bca08b4ab5db2874
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    results.tekton.dev/result: test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45
    build.appstudio.redhat.com/target_branch: main
    results.tekton.dev/childReadyForDeletion: 'true'
    build.appstudio.redhat.com/bundle: >-
      quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:docker-build-124dd76239d17a3bdb936ed26e403c9c9c4e947b
    build.appstudio.redhat.com/commit_sha: cb64992eebf5c18900d283a9bca08b4ab5db2874
  name: devfile-sample-nsj6m-show-sbom
  uid: 82872a3b-5596-45f2-84ac-d499953ad295
  creationTimestamp: '2023-08-30T19:01:30Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:build.appstudio.redhat.com/pipeline_name': {}
            'f:build.appstudio.redhat.com/bundle': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            .: {}
            'f:pipeline.tekton.dev/release': {}
            'f:tekton.dev/tags': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:tekton.dev/pipelines.minVersion': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelines.openshift.io/used-by': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:appstudio.openshift.io/application': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:appstudio.openshift.io/component': {}
            'f:pipelines.openshift.io/strategy': {}
            'f:pipelines.openshift.io/runtime': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"579cdd9f-5dcc-436c-ab9b-eb770530ff45"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-30T19:01:30Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-30T19:01:36Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:steps': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-30T19:01:36Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-30T19:02:13Z'
  namespace: test-rhtap-95-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: devfile-sample-nsj6m
      uid: 579cdd9f-5dcc-436c-ab9b-eb770530ff45
  finalizers:
    - chains.tekton.dev
  labels:
    tekton.dev/memberOf: finally
    appstudio.openshift.io/component: devfile-sample
    pipelines.openshift.io/runtime: generic
    pipelines.openshift.io/strategy: docker
    app.kubernetes.io/version: '0.1'
    tekton.dev/pipeline: docker-build
    pipelines.openshift.io/used-by: build-cloud
    app.kubernetes.io/managed-by: tekton-pipelines
    appstudio.openshift.io/application: load-app-7089021e-3585-438c-85e8-0a1ec8381a9f
    tekton.dev/task: show-sbom
    tekton.dev/pipelineTask: show-sbom
    pipelines.appstudio.openshift.io/type: build
    tekton.dev/pipelineRun: devfile-sample-nsj6m
spec:
  params:
    - name: IMAGE_URL
      value: >-
        quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: show-sbom
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:show-sbom-0.1@sha256:7db0af43dcebaeb33e34413148370e17078c30fd2fc78fb84c8941b444199f36
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
status:
  completionTime: '2023-08-30T19:01:36Z'
  conditions:
    - lastTransitionTime: '2023-08-30T19:01:36Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: devfile-sample-nsj6m-show-sbom-pod
  startTime: '2023-08-30T19:01:30Z'
  steps:
    - container: step-show-sbom
      imageID: >-
        quay.io/redhat-appstudio/cosign@sha256:c883d6f8d39148f2cea71bff4622d196d89df3e510f36c140c097b932f0dd5d5
      name: show-sbom
      terminated:
        containerID: >-
          cri-o://6b18d743bb70bdc2ccb2dbb37d80571ed897fca535ace395329032e68983abd7
        exitCode: 0
        finishedAt: '2023-08-30T19:01:36Z'
        reason: Completed
        startedAt: '2023-08-30T19:01:35Z'
  taskSpec:
    description: >-
      Shows the Software Bill of Materials (SBOM) generated for the built image
      in CyloneDX JSON format.
    params:
      - description: Fully qualified image name to show SBOM for.
        name: IMAGE_URL
        type: string
    steps:
      - env:
          - name: IMAGE_URL
            value: >-
              quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
        image: 'quay.io/redhat-appstudio/cosign:v2.1.1'
        name: show-sbom
        resources: {}
        script: |
          #!/busybox/sh
          cosign download sbom $IMAGE_URL 2>err
          RET=$?
          if [ $RET -ne 0 ]; then
            echo Failed to get SBOM >&2
            cat err >&2
          fi
          exit $RET
`

const tooBigNumTRSummYaml = `
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  annotations:
    pipeline.tekton.dev/release: b8ad1b2
    tekton.dev/tags: 'appstudio, hacbs'
    results.tekton.dev/record: >-
      test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45/records/42a9b7bd-b75e-429b-af1d-4add9543d051
    results.tekton.dev/log: >-
      test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45/logs/022a5cfc-1763-3ce6-b21c-29f0b6ea1461
    build.appstudio.redhat.com/pipeline_name: docker-build
    build.appstudio.openshift.io/repo: >-
      https://github.com/nodeshift-starters/devfile-sample?rev=cb64992eebf5c18900d283a9bca08b4ab5db2874
    chains.tekton.dev/signed: 'true'
    tekton.dev/pipelines.minVersion: 0.12.1
    results.tekton.dev/result: test-rhtap-95-tenant/results/579cdd9f-5dcc-436c-ab9b-eb770530ff45
    build.appstudio.redhat.com/target_branch: main
    results.tekton.dev/childReadyForDeletion: 'true'
    build.appstudio.redhat.com/bundle: >-
      quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:docker-build-124dd76239d17a3bdb936ed26e403c9c9c4e947b
    build.appstudio.redhat.com/commit_sha: cb64992eebf5c18900d283a9bca08b4ab5db2874
  name: devfile-sample-nsj6m-show-summary
  uid: 42a9b7bd-b75e-429b-af1d-4add9543d051
  creationTimestamp: '2023-08-30T19:01:30Z'
  generation: 1
  managedFields:
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/result': {}
            'f:build.appstudio.redhat.com/pipeline_name': {}
            'f:build.appstudio.redhat.com/bundle': {}
            'f:build.appstudio.redhat.com/target_branch': {}
            'f:build.appstudio.redhat.com/commit_sha': {}
            .: {}
            'f:pipeline.tekton.dev/release': {}
            'f:tekton.dev/tags': {}
            'f:build.appstudio.openshift.io/repo': {}
            'f:tekton.dev/pipelines.minVersion': {}
          'f:labels':
            'f:tekton.dev/task': {}
            'f:tekton.dev/pipelineTask': {}
            'f:pipelines.openshift.io/used-by': {}
            'f:pipelines.appstudio.openshift.io/type': {}
            'f:appstudio.openshift.io/application': {}
            'f:tekton.dev/memberOf': {}
            .: {}
            'f:tekton.dev/pipelineRun': {}
            'f:appstudio.openshift.io/component': {}
            'f:pipelines.openshift.io/strategy': {}
            'f:pipelines.openshift.io/runtime': {}
            'f:app.kubernetes.io/version': {}
            'f:tekton.dev/pipeline': {}
          'f:ownerReferences':
            .: {}
            'k:{"uid":"579cdd9f-5dcc-436c-ab9b-eb770530ff45"}': {}
        'f:spec':
          .: {}
          'f:params': {}
          'f:serviceAccountName': {}
          'f:taskRef':
            .: {}
            'f:kind': {}
            'f:params': {}
            'f:resolver': {}
      manager: openshift-pipelines-controller
      operation: Update
      time: '2023-08-30T19:01:30Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:chains.tekton.dev/signed': {}
          'f:finalizers':
            .: {}
            'v:"chains.tekton.dev"': {}
      manager: openshift-pipelines-chains-controller
      operation: Update
      time: '2023-08-30T19:01:36Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:status':
          .: {}
          'f:completionTime': {}
          'f:conditions': {}
          'f:podName': {}
          'f:startTime': {}
          'f:steps': {}
          'f:taskSpec':
            .: {}
            'f:description': {}
            'f:params': {}
            'f:steps': {}
      manager: openshift-pipelines-controller
      operation: Update
      subresource: status
      time: '2023-08-30T19:01:36Z'
    - apiVersion: tekton.dev/v1beta1
      fieldsType: FieldsV1
      fieldsV1:
        'f:metadata':
          'f:annotations':
            'f:results.tekton.dev/childReadyForDeletion': {}
            'f:results.tekton.dev/log': {}
            'f:results.tekton.dev/record': {}
      manager: watcher
      operation: Update
      time: '2023-08-30T19:02:15Z'
  namespace: test-rhtap-95-tenant
  ownerReferences:
    - apiVersion: tekton.dev/v1beta1
      blockOwnerDeletion: true
      controller: true
      kind: PipelineRun
      name: devfile-sample-nsj6m
      uid: 579cdd9f-5dcc-436c-ab9b-eb770530ff45
  finalizers:
    - chains.tekton.dev
  labels:
    tekton.dev/memberOf: finally
    appstudio.openshift.io/component: devfile-sample
    pipelines.openshift.io/runtime: generic
    pipelines.openshift.io/strategy: docker
    app.kubernetes.io/version: '0.1'
    tekton.dev/pipeline: docker-build
    pipelines.openshift.io/used-by: build-cloud
    app.kubernetes.io/managed-by: tekton-pipelines
    appstudio.openshift.io/application: load-app-7089021e-3585-438c-85e8-0a1ec8381a9f
    tekton.dev/task: summary
    tekton.dev/pipelineTask: show-summary
    pipelines.appstudio.openshift.io/type: build
    tekton.dev/pipelineRun: devfile-sample-nsj6m
spec:
  params:
    - name: pipelinerun-name
      value: devfile-sample-nsj6m
    - name: git-url
      value: >-
        https://github.com/nodeshift-starters/devfile-sample.git?rev=cb64992eebf5c18900d283a9bca08b4ab5db2874
    - name: image-url
      value: >-
        quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
    - name: build-task-status
      value: Succeeded
  serviceAccountName: appstudio-pipeline
  taskRef:
    kind: Task
    params:
      - name: name
        value: summary
      - name: bundle
        value: >-
          quay.io/redhat-appstudio-tekton-catalog/pull-request-builds:summary-0.1@sha256:e69f53a3991d7088d8aa2827365ab761ab7524d4269f296b4a78b0f085789d30
      - name: kind
        value: task
    resolver: bundles
  timeout: 1h0m0s
status:
  completionTime: '2023-08-30T19:01:36Z'
  conditions:
    - lastTransitionTime: '2023-08-30T19:01:36Z'
      message: All Steps have completed executing
      reason: Succeeded
      status: 'True'
      type: Succeeded
  podName: devfile-sample-nsj6m-show-summary-pod
  startTime: '2023-08-30T19:01:30Z'
  steps:
    - container: step-appstudio-summary
      imageID: >-
        registry.access.redhat.com/ubi9/ubi-minimal@sha256:dc02c6aa8199beb8ed13312d7116a94aa87b5412886bbe358845d3f0626c0f1e
      name: appstudio-summary
      terminated:
        containerID: >-
          cri-o://149f69bd03f32d573735805c98eb57f5ef7de8dd4e91a207a084bbe1f60b907d
        exitCode: 0
        finishedAt: '2023-08-30T19:01:35Z'
        reason: Completed
        startedAt: '2023-08-30T19:01:35Z'
  taskSpec:
    description: >-
      Summary Pipeline Task. Prints PipelineRun information, removes image
      repository secret used by the PipelineRun.
    params:
      - description: pipeline-run to annotate
        name: pipelinerun-name
        type: string
      - description: Git URL
        name: git-url
        type: string
      - description: Image URL
        name: image-url
        type: string
      - default: Succeeded
        description: State of build task in pipelineRun
        name: build-task-status
        type: string
    steps:
      - env:
          - name: GIT_URL
            value: >-
              https://github.com/nodeshift-starters/devfile-sample.git?rev=cb64992eebf5c18900d283a9bca08b4ab5db2874
          - name: IMAGE_URL
            value: >-
              quay.io/redhat-user-workloads-stage/test-rhtap-95-tenant/load-app-7089021e-3585-438c-85e8-0a1ec8381a9f/devfile-sample:build-a5b73-1693422002
          - name: PIPELINERUN_NAME
            value: devfile-sample-nsj6m
          - name: BUILD_TASK_STATUS
            value: Succeeded
        image: >-
          registry.access.redhat.com/ubi9/ubi-minimal:9.2-717@sha256:dc02c6aa8199beb8ed13312d7116a94aa87b5412886bbe358845d3f0626c0f1e
        name: appstudio-summary
        resources: {}
        script: |
          #!/usr/bin/env bash
          echo
          echo "Build Summary:"
          echo
          echo "Build repository: $GIT_URL"
          if [ "$BUILD_TASK_STATUS" == "Succeeded" ]; then
            echo "Generated Image is in : $IMAGE_URL"
          fi
          echo
          echo End Summary
`
