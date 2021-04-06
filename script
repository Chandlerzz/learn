$bb = ""
while(1){
$aa = (Get-Item ~/learn/scrapy/hello/hello.go).lastWriteTime
if($aa -ne $bb){
	go install
	hello
	$bb=$aa
	}
	sleep(1)
}
