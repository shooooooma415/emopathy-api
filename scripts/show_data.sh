#!/bin/bash

echo "テーブルのデータを表示します..."
echo "=================================="

echo ""
echo "users テーブルのデータ:"
echo "-------------------------"
docker-compose exec postgres psql -U postgres -d emopathy -c "SELECT * FROM users;"

echo ""
echo "users テーブルのレコード数:"
echo "-----------------------------"
docker-compose exec postgres psql -U postgres -d emopathy -c "SELECT COUNT(*) as user_count FROM users;"

echo ""
echo "groups テーブルのデータ:"
echo "--------------------------"
docker-compose exec postgres psql -U postgres -d emopathy -c "SELECT * FROM groups;"

echo ""
echo "groups テーブルのレコード数:"
echo "------------------------------"
docker-compose exec postgres psql -U postgres -d emopathy -c "SELECT COUNT(*) as group_count FROM groups;"

echo ""
echo "group_members テーブルのデータ:"
echo "--------------------------------"
docker-compose exec postgres psql -U postgres -d emopathy -c "SELECT * FROM group_members;"

echo ""
echo "group_members テーブルのレコード数:"
echo "------------------------------------"
docker-compose exec postgres psql -U postgres -d emopathy -c "SELECT COUNT(*) as member_count FROM group_members;"

echo ""
echo "user_events テーブルのデータ:"
echo "-------------------------------"
docker-compose exec postgres psql -U postgres -d emopathy -c "SELECT * FROM user_events;"

echo ""
echo "user_events テーブルのレコード数:"
echo "----------------------------------"
docker-compose exec postgres psql -U postgres -d emopathy -c "SELECT COUNT(*) as event_count FROM user_events;"
