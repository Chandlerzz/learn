#!/usr/bin/perl
use warnings;
use strict;

our %count = ("1",1);
open(my $FH,'<','source') or die $!;
while(<$FH>){
$_ =~ s/[a-zA-Z0-9=\\\?\.\[\]\(\)\{\}\$\^]/ /g;
my @words = split/ /,$_;
for my $word(@words){
	print $word,"\n";
my @keyss = keys(%count);
if(grep { $_ eq $word } @keyss){
	$count{$word} = $count{$word} +1;
        }else{
            $count{$word} = 0;
}
        
}
}
foreach my$key(sort{$count{$b}<=>$count{$a}}keys%count){my $value = $count{$key};print $key,$value,"\n";}
