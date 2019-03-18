require 'sinatra'
require 'sinatra/reloader'
require "sinatra/json"

enable :reloader
enable :sessions

root = ::File.dirname(__FILE__)
require ::File.join(root, 'src/app')

run Sinatra::Application
