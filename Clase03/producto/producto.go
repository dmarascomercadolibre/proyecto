package producto

// Definir la interfaz Product con el método Price
type Producto interface {
	Price() float64
}

// Implementar la estructura base para todos los productos
type BaseProducto struct {
	Cost float64
}

// Implementar la estructura para el producto Small
type SmallProducto struct {
	BaseProducto
}

// Implementar la estructura para el producto Medium
type MediumProducto struct {
	BaseProducto
}

// Implementar la estructura para el producto Large
type LargeProducto struct {
	BaseProducto
}

// Implementar el método Price para la estructura base
func (p *SmallProducto) Price() float64 {
	return p.Cost
}

// Implementar el método Price para el producto Medium
func (p *MediumProducto) Price() float64 {
	// Agregar un 3% al costo del producto
	return p.Cost * 1.03
}

// Implementar el método Price para el producto Large
func (p *LargeProducto) Price() float64 {
	// Agregar un 6% al costo del producto y $2500 de costo de envío
	return p.Cost*1.06 + 2500
}

// Función factory para crear productos en base al tipo y precio
func CreateProduct(productType string, cost float64) Producto {
	switch productType {
	case "Small":
		return &SmallProducto{BaseProducto{Cost: cost}}
	case "Medium":
		return &MediumProducto{BaseProducto{Cost: cost}}
	case "Large":
		return &LargeProducto{BaseProducto{Cost: cost}}
	default:
		return nil
	}
}
