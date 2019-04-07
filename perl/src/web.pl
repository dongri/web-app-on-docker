use strict;

use Mojolicious::Lite;

get '/' => sub {
  my $c = shift;
  $c->render(text => 'Hello Perl! Hello Mojolicious!');
};

app->start;
