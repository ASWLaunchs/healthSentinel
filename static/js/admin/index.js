const Index = {
    data: {},
    func: {
        init: function () {
            Date.prototype.Format = function (fmt) {
                var o = {
                    "M+": this.getMonth() + 1,
                    "d+": this.getDate(),
                    "h+": this.getHours(),
                    "m+": this.getMinutes(),
                    "s+": this.getSeconds(),
                    "q+": Math.floor((this.getMonth() + 3) / 3),
                    "S": this.getMilliseconds()
                };
                if (/(y+)/.test(fmt))
                    fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
                for (var k in o)
                    if (new RegExp("(" + k + ")").test(fmt)) fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
                return fmt;
            }

            $.get("/admin/userList", {
                queryDate: (new Date()).Format("yyyy-MM-dd")
            }, function (data, status) {
                if (status) {
                    data.msg.forEach(v => {
                        $("#list").append(`   <tr>
                    <th scope="row">${v.sno}</th>
                    <td>麻子</td>
                    <td>2022-4.10</td>
                    <td class="text-center text-success blockquote mb-0">√</td>
                </tr>`)
                    });

                    console.log(data)
                }
            })
        }

    }
}