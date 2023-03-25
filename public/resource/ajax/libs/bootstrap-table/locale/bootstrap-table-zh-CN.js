/**
 * Bootstrap Table Chinese translation
 * Author: Zhixin Wen<wenzhixin2010@gmail.com>
 */
$.fn.bootstrapTable.locales['zh-CN'] = {
  formatShowSearch: function formatShowSearch() {
    return '隐藏/显示搜索';
  },
  formatPageGo: function formatPageGo() {
    return '跳转';
  },
  formatCopyRows: function formatCopyRows() {
    return '复制行';
  },
  formatPrint: function formatPrint() {
    return '打印';
  },
  formatLoadingMessage: function formatLoadingMessage() {
    return '🕗...';
  },
  formatRecordsPerPage: function formatRecordsPerPage(pageNumber) {
    return "每页显示 ".concat(pageNumber, " 条记录");
  },
  formatShowingRows: function formatShowingRows(pageFrom, pageTo, totalRows, totalNotFiltered) {
    if (totalNotFiltered !== undefined && totalNotFiltered > 0 && totalNotFiltered > totalRows) {
      return "显示第 ".concat(pageFrom, " 到第 ").concat(pageTo, " 条记录，总共 ").concat(totalRows, " 条记录（从 ").concat(totalNotFiltered, " 总记录中过滤）");
    }
    return "显示第 ".concat(pageFrom, " 到第 ").concat(pageTo, " 条记录，总共 ").concat(totalRows, " 条记录");
  },
  formatSRPaginationPreText: function formatSRPaginationPreText() {
    return '上一页';
  },
  formatSRPaginationPageText: function formatSRPaginationPageText(page) {
    return "第".concat(page, "页");
  },
  formatSRPaginationNextText: function formatSRPaginationNextText() {
    return '下一页';
  },
  formatDetailPagination: function formatDetailPagination(totalRows) {
    return "总共 ".concat(totalRows, " 条记录");
  },
  formatClearSearch: function formatClearSearch() {
    return '清空过滤';
  },
  formatSearch: function formatSearch() {
    return '搜索';
  },
  formatNoMatches: function formatNoMatches() {
    return '没有找到匹配的记录';
  },
  formatPaginationSwitch: function formatPaginationSwitch() {
    return '隐藏/显示分页';
  },
  formatPaginationSwitchDown: function formatPaginationSwitchDown() {
    return '显示分页';
  },
  formatPaginationSwitchUp: function formatPaginationSwitchUp() {
    return '隐藏分页';
  },
  formatRefresh: function formatRefresh() {
    return '刷新';
  },
  formatToggle: function formatToggle() {
    return '切换';
  },
  formatToggleOn: function formatToggleOn() {
    return '显示卡片视图';
  },
  formatToggleOff: function formatToggleOff() {
    return '隐藏卡片视图';
  },
  formatColumns: function formatColumns() {
    return '列';
  },
  formatColumnsToggleAll: function formatColumnsToggleAll() {
    return '切换所有';
  },
  formatFullscreen: function formatFullscreen() {
    return '全屏';
  },
  formatAllRows: function formatAllRows() {
    return '所有';
  },
  formatAutoRefresh: function formatAutoRefresh() {
    return '自动刷新';
  },
  formatExport: function formatExport() {
    return '导出数据';
  },
  formatJumpTo: function formatJumpTo() {
    return '跳转';
  },
  formatAdvancedSearch: function formatAdvancedSearch() {
    return '高级搜索';
  },
  formatAdvancedCloseButton: function formatAdvancedCloseButton() {
    return '关闭';
  },
  formatFilterControlSwitch: function formatFilterControlSwitch() {
    return '隐藏/显示过滤控制';
  },
  formatFilterControlSwitchHide: function formatFilterControlSwitchHide() {
    return '隐藏过滤控制';
  },
  formatFilterControlSwitchShow: function formatFilterControlSwitchShow() {
    return '显示过滤控制';
  }
};
$.extend($.fn.bootstrapTable.defaults, $.fn.bootstrapTable.locales['zh-CN']);
