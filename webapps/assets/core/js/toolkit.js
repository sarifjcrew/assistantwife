var Toolkit = {
	
	ajaxPost : function(url, data, callback) {
		return $.ajax({
			url: url,
			type: 'POST',
			data: ko.mapping.toJSON(data),
			contentType: 'application/json charset=utf-8',
			success: function(data) {
				callback(data)
			}, 
			error: function(xhr, status, error) {
				alert(error)
			}
		})
	},
	
	convertDate: function(date) {
		return moment(date).format("DD/MM/YY")
	}
}