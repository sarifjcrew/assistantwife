var DebtModel = function() {
	var Model = {
		Template : {
			Id : "",
			Tanggal : new(Date),
			Keterangan : "",
			Debt : 0
		},
		Attr : {
			TotalHutang : ko.observable(0) 
		}
	}

	Model.data = ko.mapping.fromJS(Model.Template)
	Model.debtList = ko.observableArray([])

	var getDebt = function() {
		Toolkit.ajaxPost("/debt/get", {}, function(data){
			Model.debtList(data.Rows)
			Model.Attr.TotalHutang(data.TotalDebt)
		})
	}

	var Save = function() {
		convertParam()
		
		Toolkit.ajaxPost("/debt/save", Model.data, function(data)	{
			resetForm()
			getDebt()
		})
	}

	var editDebt = function(index) {
		return function() {
			ko.mapping.fromJS(Model.debtList()[index], Model.data)
		}
	}

	var deleteDebt = function(id) {
		return function() {
			Toolkit.ajaxPost("/debt/delete", {"id":id}, function(data){
				getDebt()
			})
		}
	}

	var convertParam = function() {
		let debt = Model.data.Debt()
		 Model.data.Debt(parseInt(debt))
	}

	var resetForm = function() {
		ko.mapping.fromJS(Model.Template, Model.data)
	}

	var init = function() {
		ko.applyBindings(DebtModel)
		getDebt()
	}

	$(init)

	return {
		Model : Model,
		Save : Save,
		editDebt : editDebt,
		deleteDebt : deleteDebt
	}
}()