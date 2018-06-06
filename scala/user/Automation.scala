// Generated by the Scala Plugin for the Protocol Buffer Compiler.
// Do not edit!
//
// Protofile syntax: PROTO2

package user

@SerialVersionUID(0L)
final case class Automation(
    ctx: scala.Option[common.Context] = None,
    accountId: scala.Option[_root_.scala.Predef.String] = None,
    id: scala.Option[_root_.scala.Predef.String] = None,
    channel: scala.Option[_root_.scala.Predef.String] = None,
    name: scala.Option[_root_.scala.Predef.String] = None,
    description: scala.Option[_root_.scala.Predef.String] = None,
    conditions: _root_.scala.collection.Seq[user.Condition] = _root_.scala.collection.Seq.empty,
    created: scala.Option[_root_.scala.Long] = None,
    modified: scala.Option[_root_.scala.Long] = None,
    state: scala.Option[_root_.scala.Predef.String] = None,
    actionType: scala.Option[_root_.scala.Predef.String] = None,
    actionData: scala.Option[_root_.scala.Predef.String] = None,
    scope: scala.Option[_root_.scala.Predef.String] = None
    ) extends scalapb.GeneratedMessage with scalapb.Message[Automation] with scalapb.lenses.Updatable[Automation] {
    @transient
    private[this] var __serializedSizeCachedValue: _root_.scala.Int = 0
    private[this] def __computeSerializedValue(): _root_.scala.Int = {
      var __size = 0
      if (ctx.isDefined) { __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(ctx.get.serializedSize) + ctx.get.serializedSize }
      if (accountId.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(2, accountId.get) }
      if (id.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(3, id.get) }
      if (channel.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(4, channel.get) }
      if (name.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(5, name.get) }
      if (description.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(6, description.get) }
      conditions.foreach(conditions => __size += 1 + _root_.com.google.protobuf.CodedOutputStream.computeUInt32SizeNoTag(conditions.serializedSize) + conditions.serializedSize)
      if (created.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(9, created.get) }
      if (modified.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeInt64Size(10, modified.get) }
      if (state.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(12, state.get) }
      if (actionType.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(13, actionType.get) }
      if (actionData.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(14, actionData.get) }
      if (scope.isDefined) { __size += _root_.com.google.protobuf.CodedOutputStream.computeStringSize(15, scope.get) }
      __size
    }
    final override def serializedSize: _root_.scala.Int = {
      var read = __serializedSizeCachedValue
      if (read == 0) {
        read = __computeSerializedValue()
        __serializedSizeCachedValue = read
      }
      read
    }
    def writeTo(`_output__`: _root_.com.google.protobuf.CodedOutputStream): _root_.scala.Unit = {
      ctx.foreach { __v =>
        _output__.writeTag(1, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      accountId.foreach { __v =>
        _output__.writeString(2, __v)
      };
      id.foreach { __v =>
        _output__.writeString(3, __v)
      };
      channel.foreach { __v =>
        _output__.writeString(4, __v)
      };
      name.foreach { __v =>
        _output__.writeString(5, __v)
      };
      description.foreach { __v =>
        _output__.writeString(6, __v)
      };
      conditions.foreach { __v =>
        _output__.writeTag(7, 2)
        _output__.writeUInt32NoTag(__v.serializedSize)
        __v.writeTo(_output__)
      };
      created.foreach { __v =>
        _output__.writeInt64(9, __v)
      };
      modified.foreach { __v =>
        _output__.writeInt64(10, __v)
      };
      state.foreach { __v =>
        _output__.writeString(12, __v)
      };
      actionType.foreach { __v =>
        _output__.writeString(13, __v)
      };
      actionData.foreach { __v =>
        _output__.writeString(14, __v)
      };
      scope.foreach { __v =>
        _output__.writeString(15, __v)
      };
    }
    def mergeFrom(`_input__`: _root_.com.google.protobuf.CodedInputStream): user.Automation = {
      var __ctx = this.ctx
      var __accountId = this.accountId
      var __id = this.id
      var __channel = this.channel
      var __name = this.name
      var __description = this.description
      val __conditions = (_root_.scala.collection.immutable.Vector.newBuilder[user.Condition] ++= this.conditions)
      var __created = this.created
      var __modified = this.modified
      var __state = this.state
      var __actionType = this.actionType
      var __actionData = this.actionData
      var __scope = this.scope
      var _done__ = false
      while (!_done__) {
        val _tag__ = _input__.readTag()
        _tag__ match {
          case 0 => _done__ = true
          case 10 =>
            __ctx = Option(_root_.scalapb.LiteParser.readMessage(_input__, __ctx.getOrElse(common.Context.defaultInstance)))
          case 18 =>
            __accountId = Option(_input__.readString())
          case 26 =>
            __id = Option(_input__.readString())
          case 34 =>
            __channel = Option(_input__.readString())
          case 42 =>
            __name = Option(_input__.readString())
          case 50 =>
            __description = Option(_input__.readString())
          case 58 =>
            __conditions += _root_.scalapb.LiteParser.readMessage(_input__, user.Condition.defaultInstance)
          case 72 =>
            __created = Option(_input__.readInt64())
          case 80 =>
            __modified = Option(_input__.readInt64())
          case 98 =>
            __state = Option(_input__.readString())
          case 106 =>
            __actionType = Option(_input__.readString())
          case 114 =>
            __actionData = Option(_input__.readString())
          case 122 =>
            __scope = Option(_input__.readString())
          case tag => _input__.skipField(tag)
        }
      }
      user.Automation(
          ctx = __ctx,
          accountId = __accountId,
          id = __id,
          channel = __channel,
          name = __name,
          description = __description,
          conditions = __conditions.result(),
          created = __created,
          modified = __modified,
          state = __state,
          actionType = __actionType,
          actionData = __actionData,
          scope = __scope
      )
    }
    def getCtx: common.Context = ctx.getOrElse(common.Context.defaultInstance)
    def clearCtx: Automation = copy(ctx = None)
    def withCtx(__v: common.Context): Automation = copy(ctx = Option(__v))
    def getAccountId: _root_.scala.Predef.String = accountId.getOrElse("")
    def clearAccountId: Automation = copy(accountId = None)
    def withAccountId(__v: _root_.scala.Predef.String): Automation = copy(accountId = Option(__v))
    def getId: _root_.scala.Predef.String = id.getOrElse("")
    def clearId: Automation = copy(id = None)
    def withId(__v: _root_.scala.Predef.String): Automation = copy(id = Option(__v))
    def getChannel: _root_.scala.Predef.String = channel.getOrElse("")
    def clearChannel: Automation = copy(channel = None)
    def withChannel(__v: _root_.scala.Predef.String): Automation = copy(channel = Option(__v))
    def getName: _root_.scala.Predef.String = name.getOrElse("")
    def clearName: Automation = copy(name = None)
    def withName(__v: _root_.scala.Predef.String): Automation = copy(name = Option(__v))
    def getDescription: _root_.scala.Predef.String = description.getOrElse("")
    def clearDescription: Automation = copy(description = None)
    def withDescription(__v: _root_.scala.Predef.String): Automation = copy(description = Option(__v))
    def clearConditions = copy(conditions = _root_.scala.collection.Seq.empty)
    def addConditions(__vs: user.Condition*): Automation = addAllConditions(__vs)
    def addAllConditions(__vs: TraversableOnce[user.Condition]): Automation = copy(conditions = conditions ++ __vs)
    def withConditions(__v: _root_.scala.collection.Seq[user.Condition]): Automation = copy(conditions = __v)
    def getCreated: _root_.scala.Long = created.getOrElse(0L)
    def clearCreated: Automation = copy(created = None)
    def withCreated(__v: _root_.scala.Long): Automation = copy(created = Option(__v))
    def getModified: _root_.scala.Long = modified.getOrElse(0L)
    def clearModified: Automation = copy(modified = None)
    def withModified(__v: _root_.scala.Long): Automation = copy(modified = Option(__v))
    def getState: _root_.scala.Predef.String = state.getOrElse("")
    def clearState: Automation = copy(state = None)
    def withState(__v: _root_.scala.Predef.String): Automation = copy(state = Option(__v))
    def getActionType: _root_.scala.Predef.String = actionType.getOrElse("")
    def clearActionType: Automation = copy(actionType = None)
    def withActionType(__v: _root_.scala.Predef.String): Automation = copy(actionType = Option(__v))
    def getActionData: _root_.scala.Predef.String = actionData.getOrElse("")
    def clearActionData: Automation = copy(actionData = None)
    def withActionData(__v: _root_.scala.Predef.String): Automation = copy(actionData = Option(__v))
    def getScope: _root_.scala.Predef.String = scope.getOrElse("")
    def clearScope: Automation = copy(scope = None)
    def withScope(__v: _root_.scala.Predef.String): Automation = copy(scope = Option(__v))
    def getFieldByNumber(__fieldNumber: _root_.scala.Int): scala.Any = {
      (__fieldNumber: @_root_.scala.unchecked) match {
        case 1 => ctx.orNull
        case 2 => accountId.orNull
        case 3 => id.orNull
        case 4 => channel.orNull
        case 5 => name.orNull
        case 6 => description.orNull
        case 7 => conditions
        case 9 => created.orNull
        case 10 => modified.orNull
        case 12 => state.orNull
        case 13 => actionType.orNull
        case 14 => actionData.orNull
        case 15 => scope.orNull
      }
    }
    def getField(__field: _root_.scalapb.descriptors.FieldDescriptor): _root_.scalapb.descriptors.PValue = {
      require(__field.containingMessage eq companion.scalaDescriptor)
      (__field.number: @_root_.scala.unchecked) match {
        case 1 => ctx.map(_.toPMessage).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 2 => accountId.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 3 => id.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 4 => channel.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 5 => name.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 6 => description.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 7 => _root_.scalapb.descriptors.PRepeated(conditions.map(_.toPMessage)(_root_.scala.collection.breakOut))
        case 9 => created.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 10 => modified.map(_root_.scalapb.descriptors.PLong).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 12 => state.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 13 => actionType.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 14 => actionData.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
        case 15 => scope.map(_root_.scalapb.descriptors.PString).getOrElse(_root_.scalapb.descriptors.PEmpty)
      }
    }
    def toProtoString: _root_.scala.Predef.String = _root_.scalapb.TextFormat.printToUnicodeString(this)
    def companion = user.Automation
}

object Automation extends scalapb.GeneratedMessageCompanion[user.Automation] {
  implicit def messageCompanion: scalapb.GeneratedMessageCompanion[user.Automation] = this
  def fromFieldsMap(__fieldsMap: scala.collection.immutable.Map[_root_.com.google.protobuf.Descriptors.FieldDescriptor, scala.Any]): user.Automation = {
    require(__fieldsMap.keys.forall(_.getContainingType() == javaDescriptor), "FieldDescriptor does not match message type.")
    val __fields = javaDescriptor.getFields
    user.Automation(
      __fieldsMap.get(__fields.get(0)).asInstanceOf[scala.Option[common.Context]],
      __fieldsMap.get(__fields.get(1)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(2)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(3)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(4)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(5)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.getOrElse(__fields.get(6), Nil).asInstanceOf[_root_.scala.collection.Seq[user.Condition]],
      __fieldsMap.get(__fields.get(7)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(8)).asInstanceOf[scala.Option[_root_.scala.Long]],
      __fieldsMap.get(__fields.get(9)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(10)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(11)).asInstanceOf[scala.Option[_root_.scala.Predef.String]],
      __fieldsMap.get(__fields.get(12)).asInstanceOf[scala.Option[_root_.scala.Predef.String]]
    )
  }
  implicit def messageReads: _root_.scalapb.descriptors.Reads[user.Automation] = _root_.scalapb.descriptors.Reads{
    case _root_.scalapb.descriptors.PMessage(__fieldsMap) =>
      require(__fieldsMap.keys.forall(_.containingMessage == scalaDescriptor), "FieldDescriptor does not match message type.")
      user.Automation(
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(1).get).flatMap(_.as[scala.Option[common.Context]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(2).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(3).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(4).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(5).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(6).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(7).get).map(_.as[_root_.scala.collection.Seq[user.Condition]]).getOrElse(_root_.scala.collection.Seq.empty),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(9).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(10).get).flatMap(_.as[scala.Option[_root_.scala.Long]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(12).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(13).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(14).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]]),
        __fieldsMap.get(scalaDescriptor.findFieldByNumber(15).get).flatMap(_.as[scala.Option[_root_.scala.Predef.String]])
      )
    case _ => throw new RuntimeException("Expected PMessage")
  }
  def javaDescriptor: _root_.com.google.protobuf.Descriptors.Descriptor = UserProto.javaDescriptor.getMessageTypes.get(40)
  def scalaDescriptor: _root_.scalapb.descriptors.Descriptor = UserProto.scalaDescriptor.messages(40)
  def messageCompanionForFieldNumber(__number: _root_.scala.Int): _root_.scalapb.GeneratedMessageCompanion[_] = {
    var __out: _root_.scalapb.GeneratedMessageCompanion[_] = null
    (__number: @_root_.scala.unchecked) match {
      case 1 => __out = common.Context
      case 7 => __out = user.Condition
    }
    __out
  }
  lazy val nestedMessagesCompanions: Seq[_root_.scalapb.GeneratedMessageCompanion[_]] = Seq.empty
  def enumCompanionForFieldNumber(__fieldNumber: _root_.scala.Int): _root_.scalapb.GeneratedEnumCompanion[_] = throw new MatchError(__fieldNumber)
  lazy val defaultInstance = user.Automation(
  )
  sealed trait State extends _root_.scalapb.GeneratedEnum {
    type EnumType = State
    def isactive: _root_.scala.Boolean = false
    def isinactive: _root_.scala.Boolean = false
    def companion: _root_.scalapb.GeneratedEnumCompanion[State] = user.Automation.State
  }
  
  object State extends _root_.scalapb.GeneratedEnumCompanion[State] {
    implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[State] = this
    @SerialVersionUID(0L)
    case object active extends State {
      val value = 0
      val index = 0
      val name = "active"
      override def isactive: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object inactive extends State {
      val value = 1
      val index = 1
      val name = "inactive"
      override def isinactive: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    final case class Unrecognized(value: _root_.scala.Int) extends State with _root_.scalapb.UnrecognizedEnum
    
    lazy val values = scala.collection.Seq(active, inactive)
    def fromValue(value: _root_.scala.Int): State = value match {
      case 0 => active
      case 1 => inactive
      case __other => Unrecognized(__other)
    }
    def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = user.Automation.javaDescriptor.getEnumTypes.get(0)
    def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = user.Automation.scalaDescriptor.enums(0)
  }
  sealed trait ActionType extends _root_.scalapb.GeneratedEnum {
    type EnumType = ActionType
    def isconversationmessage: _root_.scala.Boolean = false
    def isagentnotification: _root_.scala.Boolean = false
    def isautomationinvitemessage: _root_.scala.Boolean = false
    def companion: _root_.scalapb.GeneratedEnumCompanion[ActionType] = user.Automation.ActionType
  }
  
  object ActionType extends _root_.scalapb.GeneratedEnumCompanion[ActionType] {
    implicit def enumCompanion: _root_.scalapb.GeneratedEnumCompanion[ActionType] = this
    @SerialVersionUID(0L)
    case object conversation_message extends ActionType {
      val value = 0
      val index = 0
      val name = "conversation_message"
      override def isconversationmessage: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object agent_notification extends ActionType {
      val value = 1
      val index = 1
      val name = "agent_notification"
      override def isagentnotification: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    case object automation_invite_message extends ActionType {
      val value = 4
      val index = 2
      val name = "automation_invite_message"
      override def isautomationinvitemessage: _root_.scala.Boolean = true
    }
    
    @SerialVersionUID(0L)
    final case class Unrecognized(value: _root_.scala.Int) extends ActionType with _root_.scalapb.UnrecognizedEnum
    
    lazy val values = scala.collection.Seq(conversation_message, agent_notification, automation_invite_message)
    def fromValue(value: _root_.scala.Int): ActionType = value match {
      case 0 => conversation_message
      case 1 => agent_notification
      case 4 => automation_invite_message
      case __other => Unrecognized(__other)
    }
    def javaDescriptor: _root_.com.google.protobuf.Descriptors.EnumDescriptor = user.Automation.javaDescriptor.getEnumTypes.get(1)
    def scalaDescriptor: _root_.scalapb.descriptors.EnumDescriptor = user.Automation.scalaDescriptor.enums(1)
  }
  implicit class AutomationLens[UpperPB](_l: _root_.scalapb.lenses.Lens[UpperPB, user.Automation]) extends _root_.scalapb.lenses.ObjectLens[UpperPB, user.Automation](_l) {
    def ctx: _root_.scalapb.lenses.Lens[UpperPB, common.Context] = field(_.getCtx)((c_, f_) => c_.copy(ctx = Option(f_)))
    def optionalCtx: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[common.Context]] = field(_.ctx)((c_, f_) => c_.copy(ctx = f_))
    def accountId: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getAccountId)((c_, f_) => c_.copy(accountId = Option(f_)))
    def optionalAccountId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.accountId)((c_, f_) => c_.copy(accountId = f_))
    def id: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getId)((c_, f_) => c_.copy(id = Option(f_)))
    def optionalId: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.id)((c_, f_) => c_.copy(id = f_))
    def channel: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getChannel)((c_, f_) => c_.copy(channel = Option(f_)))
    def optionalChannel: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.channel)((c_, f_) => c_.copy(channel = f_))
    def name: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getName)((c_, f_) => c_.copy(name = Option(f_)))
    def optionalName: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.name)((c_, f_) => c_.copy(name = f_))
    def description: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getDescription)((c_, f_) => c_.copy(description = Option(f_)))
    def optionalDescription: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.description)((c_, f_) => c_.copy(description = f_))
    def conditions: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.collection.Seq[user.Condition]] = field(_.conditions)((c_, f_) => c_.copy(conditions = f_))
    def created: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getCreated)((c_, f_) => c_.copy(created = Option(f_)))
    def optionalCreated: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.created)((c_, f_) => c_.copy(created = f_))
    def modified: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Long] = field(_.getModified)((c_, f_) => c_.copy(modified = Option(f_)))
    def optionalModified: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Long]] = field(_.modified)((c_, f_) => c_.copy(modified = f_))
    def state: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getState)((c_, f_) => c_.copy(state = Option(f_)))
    def optionalState: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.state)((c_, f_) => c_.copy(state = f_))
    def actionType: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getActionType)((c_, f_) => c_.copy(actionType = Option(f_)))
    def optionalActionType: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.actionType)((c_, f_) => c_.copy(actionType = f_))
    def actionData: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getActionData)((c_, f_) => c_.copy(actionData = Option(f_)))
    def optionalActionData: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.actionData)((c_, f_) => c_.copy(actionData = f_))
    def scope: _root_.scalapb.lenses.Lens[UpperPB, _root_.scala.Predef.String] = field(_.getScope)((c_, f_) => c_.copy(scope = Option(f_)))
    def optionalScope: _root_.scalapb.lenses.Lens[UpperPB, scala.Option[_root_.scala.Predef.String]] = field(_.scope)((c_, f_) => c_.copy(scope = f_))
  }
  final val CTX_FIELD_NUMBER = 1
  final val ACCOUNT_ID_FIELD_NUMBER = 2
  final val ID_FIELD_NUMBER = 3
  final val CHANNEL_FIELD_NUMBER = 4
  final val NAME_FIELD_NUMBER = 5
  final val DESCRIPTION_FIELD_NUMBER = 6
  final val CONDITIONS_FIELD_NUMBER = 7
  final val CREATED_FIELD_NUMBER = 9
  final val MODIFIED_FIELD_NUMBER = 10
  final val STATE_FIELD_NUMBER = 12
  final val ACTION_TYPE_FIELD_NUMBER = 13
  final val ACTION_DATA_FIELD_NUMBER = 14
  final val SCOPE_FIELD_NUMBER = 15
}
