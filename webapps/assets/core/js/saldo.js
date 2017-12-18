var SaldoModel = function() {
	var Model = {
		Template : {
			Id : "",
			Saldo : 0
		}, 
		Attr : {
			saldo : ko.observable(0),
			mode : "save"
		}
	}

	Model.data = ko.mapping.fromJS(Model.Template)

	var addSaldo = function() {
		let totalSaldo = parseInt(Model.data.Saldo()) + parseInt(Model.Attr.saldo())

		Model.data.Saldo(totalSaldo)
	}

	var getSaldo = function() {
		Toolkit.ajaxPost("/saldo/get", Model.data, function(data)	{
			if(data.length > 0) {
				ko.mapping.fromJS(data[0], Model.data)
			}
			// resetForm()
			// getSaldo()
		})
	}

	var Save = function() {
		if(Model.Attr.mode == "save") {
			addSaldo()
		} else if(Model.Attr.mode == "edit") {
			Model.data.Saldo(parseInt(Model.Attr.saldo()))
		}

		Toolkit.ajaxPost("/saldo/save", Model.data, function(data)	{
			resetForm()
			getSaldo()
		})
	}

	var Edit = function() {
		Model.Attr.mode = "edit"
		Model.Attr.saldo(Model.data.Saldo())
	}

	var resetForm = function() {
		Model.Attr.saldo(0)
		Model.Attr.mode = "save"
	}

	var init = function() {
		ko.applyBindings(SaldoModel)
		getSaldo()
	}

	$(init)

	return {
		Model : Model,
		Save : Save,
		Edit : Edit
	}
}()