/*
 * Revision History:
 *     Initial: 2018/01/17        Wang RiYu
 */

 /* L语言文法定义
 程序定义:
 <main> -> program <ID> <Body>
 <Body> -> <VarIntro> <Begin>

 变量定义:
 <VarIntro> -> var <VarDef> | ε
 <VarDef> -> <IDTable>: <Type> | <IDTable>: <Type>; <VarDef>
 <IDTable> -> <ID>, <IDTable> | <ID>

 语句定义：
 <Begin> -> begin <Sentence> end
 <Sentence> -> <Execute>; <Sentence> | <Execute>
 <Execute> -> <SimpleSt> | <StructSt>
 <SimpleSt> -> <Assignment>
 <Assignment> -> <Variable>:=<Expression>
 <Variable> -> <ID>
 <StructSt> -> <Begin> | <IfSt> | <WhileSt>
 <IfSt> -> if <BoolExpress> then <Execute> | if <BoolExpress> then <Execute> else <Execute>
 <WhileSt> -> while <BoolExpress> do <Execute>

 表达式定义:
 <Expression> -> <ArithmeticExp> | <BoolExpress>
 <ArithmeticExp> -> <ArithmeticExp> + <Item> | <ArithmeticExp> - <Item> | <Item>
 <Item> -> <Item> * <Factor> | <Item> / <Factor> | <Factor>
 <Factor> -> <ArithmeticNum> | (<ArithmeticExp>)
 <ArithmeticNum> -> <ID> | <Integer> | <Real>
 <BoolExpress> -> <BoolExpress> or <BoolItem> | <BoolItem>
 <BoolItem> -> <BoolItem> and <BoolFactor> | <BoolFactor>
 <BoolFactor> -> not <BoolFactor> | <BoolValue>
 <BoolValue> -> <BoolConstant> | <ID> | (<BoolExpress>) | <RelationExpress>
 <RalationExpress> -> <ID> <Rop> <ID>
 <Rop> -> < | <= | = | > | >= | <>

 类型定义:
 <Type> -> integer | bool | real

 单词定义:
 <ID> -> <Letter> | <ID> <Letter> | <ID> <Number>
 <Integer> -> <Number> | <Integer> <Number>
 <Real> -> <Integer> | <Real> <Number>
 <BoolValue> -> true | false

 字符定义:
 <Letter> -> A│B│C│D│E│F│G│H│I│J│K│L│M│N│O│P│Q│R│S│T│U│V│W│X│Y│Z│a│b│c│d│e│f│g│h│i│j│k│l│m│n│o│p│q│r│s│t│u│v│w│x│y│z
 <Number> -> 0│1│2│3│4│5│6│7│8│9
 */

package main
