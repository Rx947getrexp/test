
1、管理后台：
1）查询文档列表
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/official_docs/list

2）新建文档
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/official_docs/add -d '{
    "type": "help",
    "name": "help-question",
    "desc": "官方文档",
    "content": "12344\n123123\n12312312\n123123sfasjdlakjdlakj\n"
}'

3）修改文档
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/official_docs/edit -d '{
    "id": 1,
    "type": "help2",
    "name": "help-question2",
    "desc": "官方文档2",
    "content": "12344\n123123\n12312312\n123123sfasjdlakjdlakj222222222\n"
}'

4）上传图片
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13001/official_docs/upload -d '{
    "files": // 跟之前一样,
    "file_type": // 跟之前一样
}'

2、应用服务接口：
1）查询文档列表
curl -X POST -H "Content-Type: application/json" -H "Lang: cn" http://localhost:13002/get_official_docs