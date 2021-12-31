# go-option-parttern
Functional Options in Go: Implementing the Options Pattern in Golang

## Reference

https://golang.cafe/blog/golang-functional-options-pattern.html
https://medium.com/star-gazers/go-options-pattern-da49185a2526
https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
http://www.inanzzz.com/index.php/post/a8fq/implementing-the-functional-options-or-options-pattern-in-golang
https://sagikazarmark.hu/blog/functional-options-on-steroids/

# Ví dụ

Một đối tượng nhà bao gồm vật liệu xây dựng, số lượng tầng và có chỗ nhà bếp hay không

```s
type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}
```

### Cách làm thông thường.

```s
func NewHouse(meterial string, hasFireplace bool, floors int) *House {
	h := &House{
		Material:     meterial,
		HasFireplace: hasFireplace,
		Floors:       floors,
	}

	return h
}
```

Khi tạo house cần truyền đầy đủ 3 tham số, đúng thứ tự. Nếu muốn thay đổi giá trị của một trường thì phải chỉ định rõ

```s
// New a house
h := NewHouse("wood", true, 2)

// changes its attribute
h.Material = "concrete"
```

### Sử dụng options parttern

Option parttern sẽ giải quyết vấn đề không cần phải truyền đầy đủ tham số, không cần thứ tự.

- Không cần truyền đầy đủ tham số

```s
func NewHouse(opts ...Option) *House {
    h := &House{}
    // do something to apply opts for h

    return h
}
```

Trong House có các trường khác nhau với loại data khác nhau, làm sao để Option có thể gán trị cho House? Option là 1 function nhận argument là một con trỏ House. Giá trị sẽ được gán cho trường của house

```s
type Option func(*House)
```

Khai báo các hàm option cho từng trường của house

```s
func WithMaterial(material string) Option {
    return func(h *House) {
        h.Material = material
    }
}
```

Một option để đặt vật liệu là bê tông

```s
func WithConcrete() Option {
    return func(h *House) {
        h.Material = "concrete"
    }
}
```

Hàm NewHouse

```s
func NewHouse(opts ...Option) *House {
    const (
		defaultFloors       = 2
		defaultHasFireplace = true
		defaultMaterial     = "wood"
	)

    h := &House{
		Material:     defaultMaterial,
		HasFireplace: defaultHasFireplace,
		Floors:       defaultFloors,
	}

    for _, opt := opts {
        opt(h)
    }

    return h
}
```

Như vậy có thể khi khai báo một House, ta có thể tuỳ ý thêm vào các option mà không cần phải đúng thứ tự. Sau này có thay đổi logic, thêm một số trường khác thì chỉ cần thêm trường vào cho House và thêm func option cho trường mới, các phần khác không ảnh hưởng.

### Link của gốc

https://www.sohamkamani.com/golang/options-pattern/
