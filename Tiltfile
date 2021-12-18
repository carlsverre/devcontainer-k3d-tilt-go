# -*- mode: Python -*-

load('ext://restart_process', 'docker_build_with_restart')

cmd_build = docker_build

if config.tilt_subcommand == "up":
    cmd_build = docker_build_with_restart

local_resource(
  'example-build',
  'go build -o build/example ./src/cmd/example/main.go',
  deps=['src'])

cmd_build(
  'example-image',
  '.',
  entrypoint=['/build/example'],
  dockerfile='deploy/Dockerfile',
  only=['./build/example'],
  live_update=[
    sync('./build/example', '/build/example'),
  ],
)

k8s_yaml('deploy/deployment.yaml')
k8s_resource('example', port_forwards=8080)
