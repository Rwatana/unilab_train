# unilab-train

posedetection.ipynb
・ポーズをカメラで認識して間接点を検出する
・必要な間接点の座標（腰、肩、肘）が作る関節角度から速度を判定する
・wifi経由でserver.goに速度のデータを送る
動かし方は上から順に実行する
プライベートインターネットで通信する場合はgit cloneのみパブリックインターネットで実行
→プライベートインターネットに切り替えて残りを実行

server.go
・posedetection.ipynbからのデータを受け取る
本番では近藤研の方にこのコードを実行してもらう
go run server.go
で実行

