// generated by stringer -type astNodeType; DO NOT EDIT

package jmespath

import "fmt"

const _astNodeType_name = "ASTEmptyASTComparatorASTCurrentNodeASTExpRefASTFunctionExpressionASTFieldASTFilterProjectionASTFlattenASTIdentityASTIndexASTIndexExpressionASTKeyValPairASTLiteralASTMultiSelectHashASTMultiSelectListASTOrExpressionASTAndExpressionASTNotExpressionASTPipeASTProjectionASTSubexpressionASTSliceASTValueProjection"

var _astNodeType_index = [...]uint16{0, 8, 21, 35, 44, 65, 73, 92, 102, 113, 121, 139, 152, 162, 180, 198, 213, 229, 245, 252, 265, 281, 289, 307}

func (i astNodeType) String() string {
	if i < 0 || i >= astNodeType(len(_astNodeType_index)-1) {
		return fmt.Sprintf("astNodeType(%d)", i)
	}
	return _astNodeType_name[_astNodeType_index[i]:_astNodeType_index[i+1]]
}
