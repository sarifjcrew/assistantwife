var DebtModel = function() {
	var Model = {
		Template : {
			Id : "",
			Tanggal : new(Date),
			Ket : "",
			Debt : 0
		}
	}

	Model.data = ko.mapping.fromJS(Model.Template)
	Model.debtList = ko.observableArray([])

	var Save = function() {
		console.log(ko.mapping.toJS(Model.data))
		// convertParam()
		
		// Toolkit.ajaxPost("/shop/save", dataAttr.data, function(data)	{
		// 	resetForm()
		// 	getShop()
		// })
	}

	var init = function() {
		ko.applyBindings(DebtModel)
		// getShop()
	}

	$(init)

	return {
		Model : Model,
		Save : Save
	}
}()