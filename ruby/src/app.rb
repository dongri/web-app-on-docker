
get '/' do
  unless session[:count]
    session[:count] = 0
  end
  session[:count] += 1
  data = {
    :message => 'hello ruby, sinatra',
    :count => session[:count]
  }
  json data
end
