package resource

import (
	"sort"

	pb "pkg/proto/common"
	"product/model"
)

// CategorySort 按照 sort 从大到小排序
type CategorySort []*pb.Category

func (a CategorySort) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a CategorySort) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a CategorySort) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Sort < a[i].Sort
}

// CategoryResource 转换产品分类树结构
func CategoryResource(list []*model.CategoryModel) []*pb.Category {
	if len(list) == 0 {
		return []*pb.Category{}
	}
	data := buildData(list)
	tree := makeTreeCore(0, data)
	return tree
}

// buildData parent_id => id => data
func buildData(list []*model.CategoryModel) map[int64]map[int64]*pb.Category {
	var data = make(map[int64]map[int64]*pb.Category)
	for _, v := range list {
		id := v.ID
		pid := v.ParentID
		if _, ok := data[pid]; !ok {
			data[pid] = make(map[int64]*pb.Category)
		}
		data[pid][id] = &pb.Category{
			Id:       int32(id),
			ParentId: int32(pid),
			Name:     v.Name,
			Sort:     v.Sort,
			Child:    make([]*pb.Category, 0),
		}
	}
	return data
}

// makeTreeCore 递归移动
func makeTreeCore(index int64, data map[int64]map[int64]*pb.Category) []*pb.Category {
	tmp := make([]*pb.Category, 0)
	for id, item := range data[index] {
		if data[id] != nil {
			item.Child = makeTreeCore(id, data)
		}
		tmp = append(tmp, item)
	}
	//重新排序
	sort.Sort(CategorySort(tmp))
	return tmp
}
