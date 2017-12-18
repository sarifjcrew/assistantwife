var IndexModel = function() {
	var dataAttr = {
		template : {
			Id : "",
			Tanggal : new(Date),
			Keterangan : "",
			Harga : 0
		}
	}

	dataAttr.data = ko.mapping.fromJS(dataAttr.template)
	dataAttr.shopList = ko.observableArray([])

	var getShop = function() {
		Toolkit.ajaxPost("/shop/get", {}, function(data){
			dataAttr.shopList(data)
		})
	}

	var saveList = function() {
		convertParam()
		
		Toolkit.ajaxPost("/shop/save", dataAttr.data, function(data)	{
			resetForm()
			getShop()
		})
	}
	
	var editShop = function(index) {
		return function(){
			ko.mapping.fromJS(dataAttr.shopList()[index], dataAttr.data)
		}
	}

	var deleteShop = function(id) {
		return function() {
			Toolkit.ajaxPost("/shop/delete", {"id":id}, function(data){
				getShop()
			})
		}
	}


	var convertParam = function() {
		harga = dataAttr.data.Harga()
		dataAttr.data.Harga(parseInt(harga))
	}

	var resetForm = function() {
		ko.mapping.fromJS(dataAttr.template, dataAttr.data)
	}

	var init = function() {
		ko.applyBindings(IndexModel)
		getShop()
	}

	$(init);

	return {
		dataAttr : dataAttr,
		getShop  : getShop,
		saveList : saveList,
		editShop : editShop,
		deleteShop : deleteShop
	}
}()